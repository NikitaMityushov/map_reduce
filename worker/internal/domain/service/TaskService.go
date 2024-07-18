package service

import (
	"github.com/NikitaMityushov/map_reduce/worker/internal/domain/model"
)

type TaskService interface {
	GetTask() (*model.Task, error)
}
