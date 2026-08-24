package database

import (
	"context"
	"log"
	"log/slog"
	"time"

	"github.com/Sahil2k07/golang-migration/internal/configs"
	gormLogger "gorm.io/gorm/logger"
)

type slogGormLogger struct {
	gormLogger.Interface
}

func newGormLogger() gormLogger.Interface {
	appConfigs := configs.GetAppConfig()

	defaultLogger := gormLogger.New(
		log.Default(),
		gormLogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  gormLogger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  !appConfigs.FileLogging,
		},
	)

	if appConfigs.JSONLogs {
		return &slogGormLogger{
			Interface: defaultLogger,
		}
	}

	return defaultLogger
}

func (l *slogGormLogger) Trace(
	ctx context.Context,
	begin time.Time,
	fc func() (string, int64),
	err error,
) {
	sql, rows := fc()

	args := []any{
		"sql", sql,
		"rows", rows,
		"elapsed", time.Since(begin).String(),
	}

	if err != nil {
		slog.ErrorContext(
			ctx,
			"GORM_ERROR",
			append(args, "error", err)...,
		)
		return
	}

	if configs.IsDevelopment {
		slog.InfoContext(ctx, "GORM", args...)
	}
}
