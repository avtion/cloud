package main

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// setupLogging 初始化日志对象
func setupLogging(level zapcore.Level) *zap.Logger {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(level)

	logger, err := config.Build()
	if err != nil {
		log.Fatal(err)
	}

	zap.ReplaceGlobals(logger)

	if _, err = zap.RedirectStdLogAt(logger, zap.InfoLevel); err != nil {
		log.Fatal(err)
	}

	return logger
}
