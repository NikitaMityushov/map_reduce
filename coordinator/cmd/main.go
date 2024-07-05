package main

import (
	"log/slog"
	"os"

	"github.com/NikitaMityushov/map_reduce/coordinator/internal/config"
)

func main() {
	// 1) config init
	cfg := config.MustLoad()

	// 2) logger init
	log := setupLogger(cfg.Env)

	log.Info("Coordinator is started", slog.Any("cfg", cfg))

	// 3) init app

	// 4) start grpc server
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
