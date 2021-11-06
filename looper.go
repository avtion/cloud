package main

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

// looperConfig 定时器配置
type looperConfig struct {
	// WeatherCron 天气定时器规则
	WeatherCron string
}

// looper 定时器对象
type looper struct {
	cfg *globalConfig
	c   *cron.Cron
}

// newLooper 新的定时器
func newLooper(cfg *globalConfig) *looper {
	looperCfg := cfg.Looper

	l := &looper{
		cfg: cfg,
		c:   cron.New(cron.WithLocation(time.Local)),
	}

	// TODO more flexible
	if looperCfg.WeatherCron != "" {
		_, errAddWeatherCron := l.c.AddFunc(looperCfg.WeatherCron, l.startPushWeather)
		if errAddWeatherCron != nil {
			log.Printf("failed to setup weather cron job, err: %v", errAddWeatherCron)
		}
	}

	return l
}

// start 运行定时器
func (l *looper) start() {
	if l == nil {
		return
	}
	l.c.Start()
}

// startPushWeather 开始推送天气情况
func (l *looper) startPushWeather() {
	yearMonthDay := time.Now().In(time.Local).Format("2006年01月02日") + getWeekString()
	header := newLarkHeader("⛱️ 今日天气 - " + yearMonthDay)

	errSendLark := newLark(l.cfg.Lark, withContentBuilders(newWeather(l.cfg.Weather))).send(header)
	if errSendLark != nil {
		log.Printf("failed to send lark msg, err: %v", errSendLark)
	}
}
