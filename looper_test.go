package main

import (
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestPushWeather(t *testing.T) {
	setupLogging(zapcore.InfoLevel)
	cfg, err := loadWorkDirConfig()
	if err != nil {
		t.Fatal(err)
	}
	l := newLooper(cfg)
	l.startPushWeather()
}
