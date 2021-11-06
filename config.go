package main

type globalConfig struct {
	// Location 时区
	Location string
	// Looper 定时器配置
	Looper *looperConfig
	// Lark 飞书推送配置
	Lark *larkConfig
	// Weather 天气 API 配置
	Weather *weatherConfig
}
