package converter

import (
	"github.com/NikitaMityushov/map_reduce/protos/gen/go/mr_rpc_v1"
	"github.com/NikitaMityushov/map_reduce/worker/internal/domain/model"
)

func ToModel(dto *mr_rpc_v1.TaskDto) *model.Task {
	return &model.Task{
		Id:       uint(dto.Id),
		TaskType: toTaskType(&dto.TaskType),
		Chunks:   dto.Chunks,
	}
}

func toTaskType(t *mr_rpc_v1.TaskType) model.TaskType {
	if t == mr_rpc_v1.TaskType_MAP.Enum() {
		return model.MAP
	} else {
		return model.REDUCE
	}
}
