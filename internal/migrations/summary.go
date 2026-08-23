package migrations

import (
	"log/slog"
	"time"
)

func printMigrationSummary(start time.Time, table, index, view, function, procedure, trigger int) {
	duration := time.Since(start)

	slog.Info(
		"Migration summary",
		"duration", duration,
		"tables", table,
		"indexes", index,
		"views", view,
		"functions", function,
		"procedures", procedure,
		"triggers", trigger,
	)
}
