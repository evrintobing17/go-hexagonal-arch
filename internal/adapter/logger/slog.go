package logger

import (
	"log/slog"
	"os"

	"github.com/evrintobing17/go-hexagonal-arch/internal/adapter/config"
	slogmulti "github.com/samber/slog-multi"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *slog.Logger

func Set(config *config.App) {
	logger = slog.New(
		slog.NewTextHandler(os.Stderr, nil),
	)

	if config.Env == "production" {
		logRotate := &lumberjack.Logger{
			Filename:   "log/app.log",
			MaxSize:    100, // in megabytes(MB)
			MaxBackups: 3,
			MaxAge:     28, // in days
			Compress:   true,
		}

		logger = slog.New(
			slogmulti.Fanout(
				slog.NewJSONHandler(logRotate, nil),
				slog.NewJSONHandler(os.Stderr, nil),
			),
		)
	}
	slog.SetDefault(logger)
}
