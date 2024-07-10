package service

import (
	"errors"

	"github.com/NikitaMityushov/map_reduce/coordinator/internal/domain/model"
	"github.com/NikitaMityushov/map_reduce/coordinator/internal/domain/repository"
)

type coordinatorServiceImpl struct {
	chunksRepository repository.ChunksRepository
}

func NewCoordinatorServiceImpl(r repository.ChunksRepository) *coordinatorServiceImpl {
	return &coordinatorServiceImpl{r}
}

func (c *coordinatorServiceImpl) GetTask() (model.Task, error) {
	chunks, err := c.chunksRepository.GetChunks()
	if err != nil {
		return model.Task{}, errors.New("chunks repository problem")
	}
	return model.Task{Id: 1, TaskType: model.MAP, TaskChunks: chunks, Status: model.CREATED}, nil
}
