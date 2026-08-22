package main

import (
	"log/slog"

	"github.com/Sahil2k07/golang-migration/internal/configs"
)

var appConfigs configs.GlobalConfig

func init() {
	appConfigs = configs.LoadConfigs()
}

func main() {
	slog.Info(appConfigs.Database.Host) // testing
}
