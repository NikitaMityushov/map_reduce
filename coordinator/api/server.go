package api

import (
	"context"
	"errors"

	domain "github.com/NikitaMityushov/map_reduce/coordinator/internal/domain/service"
	"github.com/NikitaMityushov/map_reduce/coordinator/internal/infrastructure/repository"
	infra "github.com/NikitaMityushov/map_reduce/coordinator/internal/infrastructure/service"
	rpc "github.com/NikitaMityushov/map_reduce/protos/gen/go/mr_rpc_v1"
	"google.golang.org/grpc"
)

type serverAPI struct {
	rpc.UnimplementedMapReduceServer
	coordinatorService domain.CoordinatorService
}

func RegisterServerAPI(gRPC *grpc.Server) {
	chunkRepo := repository.NewChunksRepositoryImpl()
	coorService := infra.NewCoordinatorServiceImpl(chunkRepo)

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
	return &rpc.GetTasksResponse{Tasks: tasks.TaskChunks}, nil
}
