package main

import (
	"log/slog"

	"github.com/Sahil2k07/golang-migration/internal/configs"
	"github.com/Sahil2k07/golang-migration/internal/database"
	"github.com/Sahil2k07/golang-migration/internal/migrations"
)

func init() {
	configs.LoadConfigs()
	database.ConnectPostgres()
}

func main() {
	err := migrations.RunMigrations()
	if err != nil {
		slog.Error("Migrations failed", "error", err)
		return
	}

	slog.Info("Migrations Applied Successfully")
}
