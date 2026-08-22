package configs

import (
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
	Database databaseConfig `toml:"database"`
}

func LoadConfigs() GlobalConfig {
	once.Do(func() {
		path, exists := utils.ResolveFilePath("configs/app.toml")
		if !exists {
			globalConfig = GlobalConfig{
				Database: loadDatabaseConfig(),
			}
		} else {
			_, err := toml.DecodeFile(path, &globalConfig)
			if err != nil {
				panic("failed to decode config file: " + err.Error())
			}

			IsDevelopment = true
		}
	})

	return globalConfig
}
