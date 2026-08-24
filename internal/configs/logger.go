package configs

import (
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Sahil2k07/golang-migration/internal/utils"
)

func configureLogging() {
	var output io.Writer = os.Stdout

	if globalConfig.App.FileLogging {
		logPath := getLogFilePath()
		os.MkdirAll(filepath.Dir(logPath), 0755)

		file, err := os.OpenFile(
			logPath,
			os.O_CREATE|os.O_APPEND|os.O_WRONLY,
			0666,
		)
		if err != nil {
			slog.Error("Error opening log file for writing logs", "error", err)
			panic("")
		}

		output = io.MultiWriter(os.Stdout, file)
	}

	if globalConfig.App.JSONLogs {
		handler := slog.NewJSONHandler(output, &slog.HandlerOptions{
			Level: slog.LevelInfo,
			ReplaceAttr: func(groups []string, attr slog.Attr) slog.Attr {
				if attr.Key == slog.TimeKey {
					attr.Value = slog.TimeValue(attr.Value.Time().UTC())
				}
				return attr
			},
		})

		logger := slog.New(handler)
		slog.SetDefault(logger)

		log.SetOutput(
			slog.NewLogLogger(logger.Handler(), slog.LevelInfo).Writer(),
		)
		return
	}

	log.SetOutput(output)
}

func getLogFilePath() string {
	logDir := filepath.Join(utils.ResolveProjectRoot(), "logs")

	if _, file, _, ok := runtime.Caller(6); ok {
		return filepath.Join(
			logDir,
			filepath.Base(filepath.Dir(file))+".log",
		)
	}

	if executable, err := os.Executable(); err == nil {
		name := strings.TrimSuffix(
			filepath.Base(executable),
			filepath.Ext(executable),
		)

		return filepath.Join(logDir, name+".log")
	}

	return filepath.Join(logDir, "app.log")
}
