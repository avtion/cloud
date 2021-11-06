package main

import (
	"github.com/golang-module/carbon"
)

// newCarbon 初始化 carbon
func newCarbon() carbon.Carbon {
	lang := carbon.NewLanguage()
	lang.SetLocale("zh-CN")
	return carbon.SetLanguage(lang)
}

// getWeekString 获取星期
func getWeekString() string {
	return newCarbon().Now().ToWeekString()
}
