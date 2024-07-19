package service

import (
	"context"

	"github.com/NikitaMityushov/map_reduce/protos/gen/go/mr_rpc_v1"
	grpcclient "github.com/NikitaMityushov/map_reduce/worker/internal/clients/mreduce/grpc"
	"github.com/NikitaMityushov/map_reduce/worker/internal/converter"
	"github.com/NikitaMityushov/map_reduce/worker/internal/domain/model"
)

type taskServiceImpl struct {
	cl *grpcclient.Client
}

func (cs *taskServiceImpl) GetTask() (*model.Task, error) {
	resp, err := cs.cl.Api.GetTask(context.TODO(), &mr_rpc_v1.GetTaskRequest{})
	if err != nil {
		panic(err)
	}

	return converter.ToModel(resp.Task), nil
}

func NewTaskService(
	client *grpcclient.Client,
) *taskServiceImpl {
	return &taskServiceImpl{
		cl: client,
	}
}
