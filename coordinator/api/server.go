package api

import (
	"context"

	rpc "github.com/NikitaMityushov/map_reduce/protos/gen/go/mr_rpc_v1"
	"google.golang.org/grpc"
)

type serverAPI struct {
	rpc.UnimplementedMapReduceServer
}

func RegisterServerAPI(gRPC *grpc.Server) {
	rpc.RegisterMapReduceServer(gRPC, &serverAPI{})
}

func (s *serverAPI) GetTasks(
	ctx context.Context,
	req *rpc.GetTasksRequest,
) (*rpc.GetTasksResponse, error) {
	panic("not implemented")
}
