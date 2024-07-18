package main

import (
	"log/slog"
	"os"

	"github.com/NikitaMityushov/map_reduce/worker/internal/app"
	"github.com/NikitaMityushov/map_reduce/worker/internal/config"
)

func main() {
	// 1) config init
	cfg := config.MustLoad()

	// 2) Logger init
	log := setupLogger(cfg.Env)

	// 3) init app
	application := app.Application{Log: log, Config: cfg}
	application.Run()

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
