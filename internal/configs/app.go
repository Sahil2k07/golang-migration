package configs

import (
	"os"
	"strconv"
)

type appConfig struct {
	Environment string `toml:"environment"`
	JSONLogs    bool   `toml:"json_logs"`
}

func GetAppConfig() appConfig {
	return appConfig{
		Environment: globalConfig.App.Environment,
		JSONLogs:    globalConfig.App.JSONLogs,
	}
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

func envOrDefault(value, current string) string {
	if value != "" {
		return value
	}

	return current
}
