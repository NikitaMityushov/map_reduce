package service

import model "github.com/NikitaMityushov/map_reduce/coordinator/internal/domain/model"

type CoordinatorService interface {
	GetTask() (model.Task, error)
}
