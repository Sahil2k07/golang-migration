package migrations

import (
	"fmt"
	"log/slog"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func applyMigration(tx *gorm.DB, schemaType string, migrationFiles []migrationFile) error {
	var histories []migrationHistory

	for _, file := range migrationFiles {
		if err := tx.Exec(file.SQL); err != nil {
			return fmt.Errorf("Migration failed: %v", file.FileName)
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
			return fmt.Errorf("Failed to save migration histories for %v: %v", schemaType, err)
		}
	}

	return nil
}
