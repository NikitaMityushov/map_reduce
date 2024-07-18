package app

import (
	"log/slog"

	grpcapp "github.com/NikitaMityushov/map_reduce/coordinator/internal/app/grpc"
)

type Application struct {
	GRPCSrv *grpcapp.GrpcApp
}

func New(
	log *slog.Logger,
	grpcPort int,
	chunks []string,
	nReduce int,

) *Application {
	grpcApp := grpcapp.New(log, grpcPort, chunks, nReduce)

	return &Application{
		GRPCSrv: grpcApp,
	}
}
