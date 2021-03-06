package main

import (
	"log"
	_ "time/tzdata"

	"go.uber.org/zap/zapcore"
)

func main() {
	setupLogging(zapcore.InfoLevel)
	cfg, err := loadWorkDirConfig()
	if err != nil {
		log.Printf("failed to load config, err: %v", err)
		return
	}
	newLooper(cfg).start()
}
