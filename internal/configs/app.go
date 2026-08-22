package configs

import (
	"log/slog"
	"os"
)

type appConfig struct {
	Environment string `toml:"environment"`
	JSONLogs    bool   `toml:"json_logs"`
}

func loadAppConfig() appConfig {
	return appConfig{
		Environment: os.Getenv("APP_ENV"),
		JSONLogs:    true,
	}
}

func configureLogging() {
	if globalConfig.App.JSONLogs {
		options := &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}

		handler := slog.NewJSONHandler(os.Stdout, options)
		slog.SetDefault(slog.New(handler))
	}
}
