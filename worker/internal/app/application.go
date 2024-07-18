package app

import (
	"context"
	"log/slog"
	"time"

	grpcclient "github.com/NikitaMityushov/map_reduce/worker/internal/clients/mreduce/grpc"
	"github.com/NikitaMityushov/map_reduce/worker/internal/config"
	infraService "github.com/NikitaMityushov/map_reduce/worker/internal/infrastructure/service"
)

type Application struct {
	Log    *slog.Logger
	Config *config.Config
}

func (a *Application) Run() {
	const op = "Application.Run"
	a.Log.Info(op)
	client, err := grpcclient.New(context.Background(), a.Log, "localhost:50051", time.Hour, 10)
	if err != nil {
		panic("Error with client init")
	}
	service := infraService.NewCoordinatorService(client)

	task, err := service.GetTask()
	if err != nil {
		a.Log.Error("CoordinatorService", op)
		panic("task service error")
	}

	a.Log.Info("task: ", task)

}
