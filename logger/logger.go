package logger

import (
	"log/slog"
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

func New(debug bool) (*slog.Logger, func() error, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, nil, err
	}

	logDir := filepath.Join(
		configDir, "pomodoro", "logs",
	)

	if err := os.MkdirAll(logDir, 0700); err != nil {
		return nil, nil, err
	}

	writer := &lumberjack.Logger{
		Filename:   filepath.Join(logDir, "app.log"),
		MaxSize:    5,
		MaxBackups: 3,
		MaxAge:     14,
		Compress:   true,
	}

	level := slog.LevelInfo

	if debug {
		level = slog.LevelDebug
	}

	handler := slog.NewJSONHandler(
		writer,
		&slog.HandlerOptions{
			Level:     level,
			AddSource: true,
		},
	)

	return slog.New(handler), writer.Close, nil
}
