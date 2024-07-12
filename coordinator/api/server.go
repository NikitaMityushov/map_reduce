package api

import (
	"context"
	"errors"

	domain "github.com/NikitaMityushov/map_reduce/coordinator/internal/domain/service"
	statemachine "github.com/NikitaMityushov/map_reduce/coordinator/internal/domain/state_machine"
	infra "github.com/NikitaMityushov/map_reduce/coordinator/internal/infrastructure/service"
	rpc "github.com/NikitaMityushov/map_reduce/protos/gen/go/mr_rpc_v1"
	"google.golang.org/grpc"
)

type serverAPI struct {
	rpc.UnimplementedMapReduceServer
	coordinatorService domain.CoordinatorService
}

func RegisterServerAPI(gRPC *grpc.Server, chunks []string, nReduce int) {
	state := statemachine.NewCoordinatorState(chunks, nReduce)
	coorService := infra.NewCoordinatorServiceImpl(*state)

	rpc.RegisterMapReduceServer(gRPC, &serverAPI{coordinatorService: coorService})
}

func (s *serverAPI) GetTasks(
	ctx context.Context,
	req *rpc.GetTasksRequest,
) (*rpc.GetTasksResponse, error) {
	tasks, err := s.coordinatorService.GetTask()
	if err != nil {
		return nil, errors.New("smth wrong with coordinator service")
	}
	return &rpc.GetTasksResponse{Tasks: tasks.Chunks}, nil
}
