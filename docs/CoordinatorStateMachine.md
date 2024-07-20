# Coordinator State Machine

## Overview

The Coordinator in the distributed MapReduce service is responsible for managing job submissions, task assignments, progress monitoring, and error handling. This document outlines the state machine for the Coordinator, including its states, transitions, and key functionalities.

## States

### 1. Idle
- **Description**: The Coordinator is waiting for new jobs.
- **Transitions**:
  - `Receive Job` -> `Initializing`

### 2. Initializing
- **Description**: The Coordinator initializes job details and splits the job into map tasks.
- **Transitions**:
  - `Initialization Complete` -> `Mapping`

### 3. Mapping
- **Description**: The Coordinator assigns map tasks to available workers, monitors progress, and handles worker failures.
- **Transitions**:
  - `All Map Tasks Completed` -> `Shuffling`

### 4. Shuffling
- **Description**: The Coordinator collects intermediate data from map tasks, groups data by key, and prepares reduce tasks.
- **Transitions**:
  - `Shuffling Complete` -> `Reducing`

### 5. Reducing
- **Description**: The Coordinator assigns reduce tasks to available workers, monitors progress, and handles worker failures.
- **Transitions**:
  - `All Reduce Tasks Completed` -> `Finalizing`

### 6. Finalizing
- **Description**: The Coordinator gathers final results, completes the job, and cleans up resources.
- **Transitions**:
  - `Job Complete` -> `Idle`

### 7. Error
- **Description**: The Coordinator handles any errors encountered during any phase, deciding whether to retry, fail the job, or perform any other recovery action.
- **Transitions**:
  - `Error Resolved` -> Previous state (e.g., `Initializing`, `Mapping`, etc.)
  - `Job Failed` -> `Idle`

## gRPC Services

### SubmitJob (JobRequest) -> JobResponse
- **Description**: Client submits a job to the Coordinator.
- **Action**: Triggers transition from `Idle` to `Initializing`.

### GetTask (WorkerRequest) -> TaskResponse
- **Description**: Worker requests a task (map or reduce).
- **Action**: Coordinator assigns a task based on the current state (e.g., map tasks during `Mapping` state).

### UpdateTaskStatus (TaskStatus) -> StatusResponse
- **Description**: Workers update the Coordinator with their task status (completed, failed, etc.).
- **Action**: Coordinator updates its state machine and transitions if necessary.

### Heartbeat (HeartbeatRequest) -> HeartbeatResponse
- **Description**: Workers send periodic heartbeats to indicate they are still active.
- **Action**: Coordinator detects worker failures based on missed heartbeats.

### GetJobStatus (JobStatusRequest) -> JobStatusResponse
- **Description**: Clients query the current status of a job.
- **Action**: Returns current state and progress of the job.

## State Machine Transitions

### Idle -> Initializing
- **Trigger**: Job submitted via `SubmitJob`.
- **Actions**: Validate job, split job into map tasks.

### Initializing -> Mapping
- **Trigger**: Initialization complete.
- **Actions**: Assign map tasks to workers, monitor progress.

### Mapping -> Shuffling
- **Trigger**: All map tasks completed.
- **Actions**: Collect intermediate data, group by key.

### Shuffling -> Reducing
- **Trigger**: Shuffling complete.
- **Actions**: Assign reduce tasks to workers, monitor progress.

### Reducing -> Finalizing
- **Trigger**: All reduce tasks completed.
- **Actions**: Gather final results, clean up resources.

### Finalizing -> Idle
- **Trigger**: Job complete.
- **Actions**: Notify client, reset Coordinator state.

### Any State -> Error
- **Trigger**: Error encountered.
- **Actions**: Log error, decide recovery action.

### Error -> Previous State
- **Trigger**: Error resolved.
- **Actions**: Resume normal operations.

### Error -> Idle
- **Trigger**: Job failed.
- **Actions**: Notify client, reset Coordinator state.

## Coordinator Struct

```go
type Coordinator struct {
    state      State
    stateMutex sync.Mutex
    tasks      map[string]*pb.TaskStatus // Store task statuses
    workers    map[string]*WorkerInfo    // Worker information
    job        *JobInfo                  // Current job information
}

type WorkerInfo struct {
    id       string
    address  string
    lastSeen time.Time
}

type JobInfo struct {
    jobId      string
    mapTasks   []*pb.Task
    reduceTasks []*pb.Task
}
```

## State Transition Methods

### setState

```go
func (c *Coordinator) setState(newState State) {
    c.stateMutex.Lock()
    defer c.stateMutex.Unlock()
    c.state = newState
}
```

### getState

```go
func (c *Coordinator) getState() State {
    c.stateMutex.Lock()
    defer c.stateMutex.Unlock()
    return c.state
}
```

## gRPC Handlers

### handleJobSubmission

```go
func (c *Coordinator) handleJobSubmission(ctx context.Context, req *pb.JobRequest) (*pb.JobResponse, error) {
    c.setState(Initializing)
    c.job = &JobInfo{
        jobId: req.JobId,
        mapTasks: initializeMapTasks(req),
    }
    c.setState(Mapping)
    c.assignTasks()
    return &pb.JobResponse{JobId: c.job.jobId}, nil
}
```

### handleTaskUpdate

```go
func (c *Coordinator) handleTaskUpdate(ctx context.Context, req *pb.TaskStatus) (*pb.StatusResponse, error) {
    c.stateMutex.Lock()
    c.tasks[req.TaskId] = req
    c.stateMutex.Unlock()

    switch c.getState() {
    case Mapping:
        if allMapTasksCompleted(c.tasks) {
            c.setState(Shuffling)
            c.setState(Reducing)
            c.assignTasks()
        }
    case Reducing:
        if allReduceTasksCompleted(c.tasks) {
            c.setState(Finalizing)
            c.setState(Idle)
        }
    }

    return &pb.StatusResponse{Status: "ok"}, nil
}
```

### handleHeartbeat

```go
func (c *Coordinator) handleHeartbeat(ctx context.Context, req *pb.HeartbeatRequest) (*pb.HeartbeatResponse, error) {
    c.stateMutex.Lock()
    worker, exists := c.workers[req.WorkerId]
    if !exists {
        worker = &WorkerInfo{id: req.WorkerId, address: req.Address}
        c.workers[req.WorkerId] = worker
    }
    worker.lastSeen = time.Now()
    c.stateMutex.Unlock()

    return &pb.HeartbeatResponse{Status: "ok"}, nil
}
```

### getJobStatus

```go
func (c *Coordinator) getJobStatus(ctx context.Context, req *pb.JobStatusRequest) (*pb.JobStatusResponse, error) {
    return &pb.JobStatusResponse{State: int32(c.getState()), JobId: c.job.jobId}, nil
}
```

## Helper Functions

### allMapTasksCompleted

```go
func allMapTasksCompleted(tasks map[string]*pb.TaskStatus) bool {
    for _, status := range tasks {
        if status.Type == pb.TaskType_MAP && status.Status != pb.Status_COMPLETED {
            return false
        }
    }
    return true
}
```

### allReduceTasksCompleted

```go
func allReduceTasksCompleted(tasks map[string]*pb.TaskStatus) bool {
    for _, status := range tasks {
        if status.Type == pb.TaskType_REDUCE && status.Status != pb.Status_COMPLETED {
            return false
        }
    }
    return true
}
```

### initializeMapTasks

```go
func initializeMapTasks(req *pb.JobRequest) []*pb.Task {
    // Split the job into map tasks
    return []*pb.Task{}
}
```

### assignTasks

```go
func (c *Coordinator) assignTasks() {
    // Assign tasks to available workers
}
```
