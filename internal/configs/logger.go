package configs

import (
	"log"
	"log/slog"
	"os"
)

func configureLogging() {
	if globalConfig.App.JSONLogs {
		options := &slog.HandlerOptions{
			Level: slog.LevelInfo,
			ReplaceAttr: func(groups []string, attr slog.Attr) slog.Attr {
				if attr.Key == slog.TimeKey {
					return slog.Attr{
						Key:   slog.TimeKey,
						Value: slog.TimeValue(attr.Value.Time().UTC()),
					}
				}

				return attr
			},
		}

		handler := slog.NewJSONHandler(os.Stdout, options)
		logger := slog.New(handler)

		slog.SetDefault(logger)

		log.SetOutput(
			slog.NewLogLogger(logger.Handler(), slog.LevelInfo).Writer(),
		)
	}
}
