package model

type Task struct {
	Id       uint
	TaskType TaskType
	Chunks   []string
}

type TaskType uint8

const (
	MAP TaskType = iota
	REDUCE
)

type Status uint8

const (
	CREATED Status = iota
	IN_PROCCESS
	DONE
)
