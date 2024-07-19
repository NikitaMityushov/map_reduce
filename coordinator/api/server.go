package api

import (
	"context"
	"errors"

	domain "github.com/NikitaMityushov/map_reduce/coordinator/internal/domain/service"
	statemachine "github.com/NikitaMityushov/map_reduce/coordinator/internal/domain/state_machine"
	infra "github.com/NikitaMityushov/map_reduce/coordinator/internal/infrastructure/service"
	"github.com/NikitaMityushov/map_reduce/coordinator/internal/mapper"
	rpc "github.com/NikitaMityushov/map_reduce/protos/gen/go/mr_rpc_v1"
	"google.golang.org/grpc"
)

type serverAPI struct {
	rpc.UnimplementedMapReduceServer
	coordinatorService domain.CoordinatorService
}

func RegisterServerAPI(gRPC *grpc.Server, chunks []string, nReduce int) {
	state := statemachine.InitCoordinatorState(chunks, nReduce)
	coorService := infra.NewCoordinatorServiceImpl(state)

	rpc.RegisterMapReduceServer(gRPC, &serverAPI{coordinatorService: coorService})
}

func (s *serverAPI) GetTask(
	ctx context.Context,
	req *rpc.GetTaskRequest,
) (*rpc.GetTaskResponse, error) {
	task, err := s.coordinatorService.GetTask()
	if err != nil {
		return nil, errors.New("smth wrong with coordinator service")
	}
	return &rpc.GetTaskResponse{Task: mapper.ToTaskDto(task)}, nil
}
