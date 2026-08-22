package database

import (
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/Sahil2k07/golang-migration/internal/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectPostgres() {
	var gormLog gormLogger.Interface

	if !configs.IsDevelopment {
		gormLog = gormLogger.Default.LogMode(gormLogger.Error)
	} else {
		stdLogger := log.New(os.Stdout, "\r\n", log.LstdFlags)
		gormLog = gormLogger.New(stdLogger, gormLogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  gormLogger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  configs.IsDevelopment,
		})
	}

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
