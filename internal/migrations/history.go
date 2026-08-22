package migrations

import (
	"fmt"

	"github.com/Sahil2k07/golang-migration/internal/database"
)

type migrationHistory struct {
}

func getMigrationHistory() ([]migrationHistory, error) {
	if err := ensureMigrationHistoryExists(); err != nil {
		return nil, err
	}

	var histories []migrationHistory

	err := database.DB.Table("migration.migration_history").Find(&histories).Error
	if err != nil {
		return nil, fmt.Errorf("Migration History error: %v", err)
	}

	return histories, nil
}

func ensureMigrationHistoryExists() error {
	sql := `
		CREATE SCHEMA IF NOT EXISTS migration;

		CREATE TABLE IF NOT EXISTS migration.migration_history (
			id         INT PRIMARY KEY,
			file_name  VARCHAR(100) NOT NULL,
			file_path  VARCHAR(200) UNIQUE NOT NULL,
			file_hash  VARCHAR(100) NOT NULL,
			file_type  VARCHAR(50) NOT NULL,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW()
		);
	`

	if err := database.DB.Exec(sql).Error; err != nil {
		return fmt.Errorf("History table error: %v", err)
	}

	return nil
}
