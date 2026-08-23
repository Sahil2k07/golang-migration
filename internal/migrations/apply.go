package migrations

import (
	"fmt"
	"log/slog"
	"time"

	"gorm.io/gorm"
)

func applyMigration(tx *gorm.DB, schemaType string, migrationFiles []migrationFile) error {
	var histories []migrationHistory

	for _, file := range migrationFiles {
		if err := tx.Exec(file.SQL).Error; err != nil {
			return fmt.Errorf("Migration failed %v: %w", file.FileName, err)
		}

		history := migrationHistory{
			FileName:  file.FileName,
			FilePath:  file.FilePath,
			FileHash:  file.FileHash,
			FileType:  schemaType,
			UpdatedAt: time.Now(),
		}

		histories = append(histories, history)

		slog.Info("Migration Applied", "file", file.FileName)
	}

	return saveMigrationHistory(tx, schemaType, histories)
}
