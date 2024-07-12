package repository

import (
	domainErrors "github.com/NikitaMityushov/map_reduce/coordinator/internal/domain/errors"
	statemachine "github.com/NikitaMityushov/map_reduce/coordinator/internal/domain/state_machine"
)

type StateRepository interface {
	SaveState(s statemachine.CoordinatorState) domainErrors.CoordinatorError
	RestoreState() (statemachine.CoordinatorState, domainErrors.CoordinatorError)
}
