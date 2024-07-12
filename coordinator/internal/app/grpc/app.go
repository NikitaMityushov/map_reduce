package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/NikitaMityushov/map_reduce/coordinator/api"
	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(
	log *slog.Logger,
	port int,
	chunks []string,
	nReduce int,
) *App {
	gRPCServer := grpc.NewServer()

	api.RegisterServerAPI(gRPCServer, chunks, nReduce)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(slog.String("op", op), slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Coordinator GRPC server got started", slog.String("addr", l.Addr().String()))
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op), slog.Int("port", a.port)).Info("Coordinator GRPC server stopped")
	a.gRPCServer.GracefulStop()
}
