package database

import (
	"log/slog"

	"github.com/Sahil2k07/golang-migration/internal/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres() {
	gormLog := newGormLogger()

	postgresDSN := configs.GetDbString()

	db, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{
		Logger: gormLog,
	})
	if err != nil {
		slog.Error("failed to connect to DB", "error", err)
		panic("Database was not found")
	}

	DB = db
}
