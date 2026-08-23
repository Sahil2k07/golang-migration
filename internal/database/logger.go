package database

import (
	"context"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/Sahil2k07/golang-migration/internal/configs"
	gormLogger "gorm.io/gorm/logger"
)

type slogGormLogger struct {
	logger *slog.Logger
}

func newGormLogger() gormLogger.Interface {
	if configs.GetAppConfig().JSONLogs {
		return &slogGormLogger{
			logger: slog.Default(),
		}
	}

	return gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		gormLogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  gormLogger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
}

func (l *slogGormLogger) LogMode(gormLogger.LogLevel) gormLogger.Interface {
	return l
}

func (l *slogGormLogger) Info(ctx context.Context, msg string, data ...any) {
	l.logger.InfoContext(ctx, msg, data...)
}

func (l *slogGormLogger) Warn(ctx context.Context, msg string, data ...any) {
	l.logger.WarnContext(ctx, msg, data...)
}

func (l *slogGormLogger) Error(ctx context.Context, msg string, data ...any) {
	l.logger.ErrorContext(ctx, msg, data...)
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
		l.logger.ErrorContext(ctx, "gorm query error", append(args, "error", err)...)
		return
	}

	if configs.IsDevelopment {
		l.logger.InfoContext(ctx, "gorm query", args...)
	}
}
