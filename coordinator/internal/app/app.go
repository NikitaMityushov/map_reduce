package app

import (
	"log/slog"

	grpcapp "github.com/NikitaMityushov/map_reduce/coordinator/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	chunks []string,
	nReduce int,

) *App {
	grpcApp := grpcapp.New(log, grpcPort, chunks, nReduce)

	return &App{
		GRPCSrv: grpcApp,
	}
}
