package main

import (
	"log/slog"

	"github.com/Sahil2k07/golang-migration/internal/configs"
	"github.com/Sahil2k07/golang-migration/internal/database"
)

var appConfigs configs.GlobalConfig

func init() {
	appConfigs = configs.LoadConfigs()
	database.ConnectPostgres()
}

func main() {
	slog.Info(appConfigs.Database.Host) // testing

	slog.Info("checking env", "isdevelopment", configs.IsDevelopment)
}
