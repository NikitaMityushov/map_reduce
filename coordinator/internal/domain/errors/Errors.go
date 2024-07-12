package errors

type CoordinatorError interface {
	Message() string
}

type TaskNotFoundError struct{}

func (e TaskNotFoundError) Message() string {
	return "there are no tasks in coordinator"
}

type InternalServerError struct {
	Info string
}

func (e InternalServerError) Message() string {
	return e.Info
}
