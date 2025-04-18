package app

import (
	"flag"
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

var (
	env     = flag.String("env", "prod", "server environment")
	logFile = flag.String("log-file", "", "file to write logs (optional)")
)

func setupLogger() *slog.Logger {
	flag.Parse()

	var writers []io.Writer

	if *env == "staging" {
		writers = append(writers, os.Stdout)
	}

	if *logFile != "" {
		dir := filepath.Dir(*logFile)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			panic("cannot create log directory: " + err.Error())
		}
		f, err := os.OpenFile(
			*logFile,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0o644,
		)
		if err != nil {
			panic("cannot open log file: " + err.Error())
		}
		writers = append(writers, f)
	}

	mw := io.MultiWriter(writers...)
	handler := slog.NewJSONHandler(mw, nil)

	logger := slog.New(handler).
		With(
			slog.String("service", "pvz"),
			slog.String("env", *env),
		)

	return logger
}
