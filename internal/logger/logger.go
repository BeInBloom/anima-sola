package logger

import (
	"log/slog"
	"os"
)

func New(env string) *slog.Logger {
	switch env {
	case "dev":
		return devLogger()
	default:
		panic("unknown environment")
	}
}

func devLogger() *slog.Logger {
	log := slog.New(slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		},
	))

	return log
}
