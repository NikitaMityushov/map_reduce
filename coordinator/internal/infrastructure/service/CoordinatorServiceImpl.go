package service

import (
	"sync"

	"github.com/NikitaMityushov/map_reduce/coordinator/internal/domain/errors"
	"github.com/NikitaMityushov/map_reduce/coordinator/internal/domain/model"
	statemachine "github.com/NikitaMityushov/map_reduce/coordinator/internal/domain/state_machine"
)

type coordinatorServiceImpl struct {
	mu               *sync.Mutex
	CoordinatorState statemachine.CoordinatorState
}

func NewCoordinatorServiceImpl(st statemachine.CoordinatorState) *coordinatorServiceImpl {

	return &coordinatorServiceImpl{
		mu:               new(sync.Mutex),
		CoordinatorState: st,
	}
}

func (c *coordinatorServiceImpl) GetTask() (model.Task, errors.CoordinatorError) {
	c.mu.Lock()
	defer c.mu.Unlock()

	task, err := c.CoordinatorState.TaskRequested()
	if err != nil {
		return model.Task{}, err
	} else {
		return task, nil
	}
}
