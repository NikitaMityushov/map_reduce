package mapper

import (
	"github.com/NikitaMityushov/map_reduce/coordinator/internal/domain/model"
	"github.com/NikitaMityushov/map_reduce/protos/gen/go/mr_rpc_v1"
)

func ToTaskDto(task model.Task) *mr_rpc_v1.TaskDto {
	var tType mr_rpc_v1.TaskType
	if task.TaskType == model.MAP {
		tType = mr_rpc_v1.TaskType_MAP
	} else {
		tType = mr_rpc_v1.TaskType_REDUCE
	}

	return &mr_rpc_v1.TaskDto{
		Id:       int64(task.Id),
		TaskType: tType,
		Chunks:   task.Chunks,
	}
}
