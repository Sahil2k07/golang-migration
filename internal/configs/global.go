package configs

import (
	"log/slog"
	"strings"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/Sahil2k07/golang-migration/internal/utils"
)

var (
	IsDevelopment bool = false
	globalConfig  GlobalConfig
	once          sync.Once
)

type GlobalConfig struct {
	App      appConfig      `toml:"app"`
	Database databaseConfig `toml:"database"`
}

func LoadConfigs() GlobalConfig {
	once.Do(func() {
		globalConfig = GlobalConfig{
			App:      loadAppConfig(),
			Database: loadDatabaseConfig(),
		}

		path, exists := utils.ResolveFilePath("configs/app.toml")
		if globalConfig.App.Environment == "" && exists {
			_, err := toml.DecodeFile(path, &globalConfig)
			if err != nil {
				slog.Error("failed to decode app.toml", "error", err)
				panic("Failed to load development configurations")
			}
		}

		IsDevelopment = strings.EqualFold(
			globalConfig.App.Environment,
			"development",
		)

		configureLogging()
	})

	return globalConfig
}
