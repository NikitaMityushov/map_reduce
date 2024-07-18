package service

import (
	"context"

	"github.com/NikitaMityushov/map_reduce/protos/gen/go/mr_rpc_v1"
	grpcclient "github.com/NikitaMityushov/map_reduce/worker/internal/clients/mreduce/grpc"
	"github.com/NikitaMityushov/map_reduce/worker/internal/domain/model"
)

type taskServiceImpl struct {
	cl *grpcclient.Client
}

func (cs *taskServiceImpl) GetTask() (*model.Task, error) {
	resp, err := cs.cl.Api.GetTasks(context.TODO(), &mr_rpc_v1.GetTasksRequest{})
	if err != nil {
		panic("tasks not found")
	}

	return &model.Task{Id: 1, TaskType: model.MAP, Chunks: resp.Tasks, Status: model.IN_PROCCESS}, nil // todo
}

func NewCoordinatorService(
	client *grpcclient.Client,
) *taskServiceImpl {
	return &taskServiceImpl{
		cl: client,
	}
}
