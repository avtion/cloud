package main

import (
	"encoding/json"
	"os"
	"testing"
)

func Test_getLocalWeather(t *testing.T) {
	cfg, err := loadWorkDirConfig()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := newWeather(cfg.Weather).getLocalWeather()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
	data, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}
	_ = os.WriteFile("weather_example.json", data, 0600)
}
