package service

import (
	domainErrors "github.com/NikitaMityushov/map_reduce/coordinator/internal/domain/errors"
	domainModels "github.com/NikitaMityushov/map_reduce/coordinator/internal/domain/model"
)

type CoordinatorService interface {
	GetTask() (domainModels.Task, domainErrors.CoordinatorError)
}
