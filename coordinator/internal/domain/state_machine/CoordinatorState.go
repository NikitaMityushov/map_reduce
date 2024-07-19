package statemachine

import (
	domainErrors "github.com/NikitaMityushov/map_reduce/coordinator/internal/domain/errors"
	domainModels "github.com/NikitaMityushov/map_reduce/coordinator/internal/domain/model"
)

// State machine
type CoordinatorState struct {
	MapTasks    []domainModels.Task
	ReduceTasks []domainModels.Task
	InProgress  []domainModels.Task
	Done        []domainModels.Task
}

// Constructor
func InitCoordinatorState(chunks []string, nReduce int) CoordinatorState {
	mapTasks := make([]domainModels.Task, 0, len(chunks))
	var curTaskId uint = 0
	for _, ch := range chunks {
		task := domainModels.Task{Id: curTaskId, TaskType: domainModels.MAP, Chunks: []string{ch}, Status: domainModels.CREATED}
		mapTasks = append(mapTasks, task)
		curTaskId++
	}

	return CoordinatorState{
		MapTasks: mapTasks,
	}
}

// Transitions
func (s *CoordinatorState) TaskRequested() (domainModels.Task, domainErrors.CoordinatorError) {
	if len(s.MapTasks) == 0 {
		if len(s.ReduceTasks) == 0 {
			return domainModels.Task{}, domainErrors.TaskNotFoundError{}
		} else {
			index := len(s.ReduceTasks) - 1
			el := (s.ReduceTasks[index])
			s.ReduceTasks = s.ReduceTasks[:index]
			return el, nil
		}
	} else {
		index := len(s.MapTasks) - 1
		el := (s.MapTasks[index])
		s.MapTasks = s.MapTasks[:index]
		return el, nil
	}
}
