package configs

import (
	"os"

	"github.com/Sahil2k07/golang-migration/internal/utils"
)

type appConfig struct {
	Environment string `toml:"environment"`
	JSONLogs    bool   `toml:"json_logs"`
	FileLogging bool   `toml:"file_logging"`
}

func GetAppConfig() appConfig {
	return appConfig{
		Environment: globalConfig.App.Environment,
		JSONLogs:    globalConfig.App.JSONLogs,
		FileLogging: globalConfig.App.FileLogging,
	}
}

func loadAppConfig(config appConfig) appConfig {
	return appConfig{
		Environment: envOrDefault(
			os.Getenv("APP_ENV"),
			config.Environment,
		),
		JSONLogs:    utils.StringToBool(os.Getenv("JSON_LOGS"), globalConfig.App.JSONLogs),
		FileLogging: utils.StringToBool(os.Getenv("FILE_LOGGING"), globalConfig.App.FileLogging),
	}
}

func envOrDefault(value, current string) string {
	if value != "" {
		return value
	}

	return current
}
