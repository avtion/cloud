package main

import (
	"testing"
)

func Test_loadWorkDirConfig(t *testing.T) {
	cfg, err := loadWorkDirConfig()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cfg)
}
