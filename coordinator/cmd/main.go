package main

import (
	"log/slog"
	"os"

	"github.com/NikitaMityushov/map_reduce/coordinator/internal/app"
	"github.com/NikitaMityushov/map_reduce/coordinator/internal/config"
)

func main() {
	// 1) config init
	cfg := config.MustLoad()

	// 2) logger init
	log := setupLogger(cfg.Env)

	// 3) init app
	application := app.New(log, cfg.GRPC.Port)

	// 4) start grpc server
	application.GRPCSrv.MustRun()

	const op = "main"
	log.With(slog.String("op", op)).Info("Coordinator is started", slog.Any("cfg", cfg))
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "dev":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "prod":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
