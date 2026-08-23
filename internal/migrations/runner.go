package migrations

import (
	"log/slog"
	"time"

	"github.com/Sahil2k07/golang-migration/internal/database"
	"gorm.io/gorm"
)

func RunMigrations() error {
	start := time.Now()

	slog.Info("Reading Files")

	histories, err := getMigrationHistory()
	if err != nil {
		return err
	}

	tableFiles, err := getMigrationFiles("migrations/tables", true, histories)
	if err != nil {
		return err
	}

	indexFiles, err := getMigrationFiles("migrations/indexes", true, histories)
	if err != nil {
		return err
	}

	viewFiles, err := getMigrationFiles("migrations/views", false, histories)
	if err != nil {
		return err
	}

	functionFiles, err := getMigrationFiles("migrations/functions", false, histories)
	if err != nil {
		return err
	}

	procedureFiles, err := getMigrationFiles("migrations/procedures", false, histories)
	if err != nil {
		return err
	}

	triggerFiles, err := getMigrationFiles("migrations/triggers", false, histories)
	if err != nil {
		return err
	}

	slog.Info("Starting Applying Migrations")

	err = database.DB.Transaction(func(tx *gorm.DB) error {
		if err := applyMigration(tx, "TABLE", tableFiles); err != nil {
			return err
		}

		if err := applyMigration(tx, "INDEX", indexFiles); err != nil {
			return err
		}

		if err := applyMigration(tx, "VIEW", viewFiles); err != nil {
			return err
		}

		if err := applyMigration(tx, "FUNCTION", functionFiles); err != nil {
			return err
		}

		if err := applyMigration(tx, "PROCEDURE", procedureFiles); err != nil {
			return err
		}

		if err := applyMigration(tx, "TRIGGER", triggerFiles); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	printMigrationSummary(
		start,
		len(tableFiles),
		len(indexFiles),
		len(viewFiles),
		len(functionFiles),
		len(procedureFiles),
		len(triggerFiles),
	)

	return nil
}
