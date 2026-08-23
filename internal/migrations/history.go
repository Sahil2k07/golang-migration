package migrations

import (
	"fmt"
	"time"

	"github.com/Sahil2k07/golang-migration/internal/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type migrationHistory struct {
	ID        int
	FileName  string
	FilePath  string
	FileHash  string
	FileType  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func getMigrationHistory() ([]migrationHistory, error) {
	if err := ensureMigrationHistoryExists(); err != nil {
		return nil, err
	}

	var histories []migrationHistory

	err := database.DB.Table("migration.migration_history").Find(&histories).Error
	if err != nil {
		return nil, fmt.Errorf("Migration History error: %w", err)
	}

	return histories, nil
}

func saveMigrationHistory(tx *gorm.DB, schemaType string, histories []migrationHistory) error {
	if len(histories) > 0 {
		if err := tx.Table("migration.migration_history").
			Clauses(clause.OnConflict{
				Columns: []clause.Column{{Name: "file_path"}},
				DoUpdates: clause.AssignmentColumns([]string{
					"updated_at",
					"file_type",
					"file_hash",
				}),
			}).
			Create(&histories).Error; err != nil {
			return fmt.Errorf("Failed to save migration histories for %v: %w", schemaType, err)
		}
	}

	return nil
}

func ensureMigrationHistoryExists() error {
	sql := `
		CREATE SCHEMA IF NOT EXISTS migration;

		CREATE TABLE IF NOT EXISTS migration.migration_history (
			id         BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
			file_name  VARCHAR(100) NOT NULL,
			file_path  VARCHAR(200) UNIQUE NOT NULL,
			file_hash  VARCHAR(100) NOT NULL,
			file_type  VARCHAR(50) NOT NULL,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW()
		);
	`

	if err := database.DB.Exec(sql).Error; err != nil {
		return fmt.Errorf("History table error: %w", err)
	}

	return nil
}
