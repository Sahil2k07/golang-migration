package configs

import (
	"log/slog"
	"os"
	"strconv"
)

type appConfig struct {
	Environment string `toml:"environment"`
	JSONLogs    bool   `toml:"json_logs"`
}

func loadAppConfig(config appConfig) appConfig {
	return appConfig{
		Environment: envOrDefault(
			os.Getenv("APP_ENV"),
			config.Environment,
		),
		JSONLogs: func() bool {
			value := os.Getenv("JSON_LOGS")
			if value == "" {
				return config.JSONLogs
			}

			parsed, err := strconv.ParseBool(value)
			if err != nil {
				return config.JSONLogs
			}

			return parsed
		}(),
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

func envOrDefault(value, current string) string {
	if value != "" {
		return value
	}

	return current
}
