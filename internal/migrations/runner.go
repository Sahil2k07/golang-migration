package migrations

import "log/slog"

func RunMigrations() error {
	histories, err := getMigrationHistory()
	if err != nil {
		return err
	}

	tableFiles, err := getMigrationFiles("migrations", true, histories)
	if err != nil {
		return err
	}

	for _, file := range tableFiles {
		slog.Info("file", "name", file.FileName, "sql", file.SQL)
	}

	return nil
}
