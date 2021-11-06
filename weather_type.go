package main

// weatherConfig 墨迹天气 API 配置
type weatherConfig struct {
	// CityID 墨迹天气城市 ID,
	// example: 广东省广州市天河区-285121
	CityID string

	// AppKey 阿里云接口调用配置
	AppKey string

	// AppSecret 阿里云接口调用配置
	AppSecret string

	// AppCode 阿里云接口调用配置
	AppCode string
}

type (
	// weatherResp 天气预报响应
	weatherResp struct {
		Code int          `json:"code"`
		Msg  string       `json:"msg"`
		Data *weatherData `json:"data"`
	}

	// weatherData 天气数据
	weatherData struct {
		// 城市信息
		City *weatherCity `json:"city"`
		// 天气预报
		Forecast []*weatherForecast `json:"forecast"`
	}

	// weatherCity 城市信息
	weatherCity struct {
		// 城市 ID
		CityId int `json:"cityId"`
		// 国家名称
		Counname string `json:"counname"`
		// 区级名称
		Name string `json:"name"`
		// 市级名称
		Pname string `json:"pname"`
	}

	// weatherForecast 预报内容
	weatherForecast struct {
		// 白天天气现象
		ConditionDay string `json:"conditionDay"`
		// 白天天气 id
		ConditionIdDay string `json:"conditionIdDay"`
		// 夜间天气id
		ConditionIdNight string `json:"conditionIdNight"`
		// 夜间天气现象
		ConditionNight string `json:"conditionNight"`
		// 预报日期 yyyy-MM-dd
		PredictDate string `json:"predictDate"`
		// 白天温度
		TempDay string `json:"tempDay"`
		// 夜间温度
		TempNight string `json:"tempNight"`
		// 更新时间 yyyy-MM-dd HH:mm:ss
		Updatetime string `json:"updatetime"`
		// 白天风向角度
		WindDirDay string `json:"windDirDay"`
		// 夜间风向角度
		WindDirNight string `json:"windDirNight"`
		// 白天风级
		WindLevelDay string `json:"windLevelDay"`
		// 夜间风级
		WindLevelNight string `json:"windLevelNight"`
	}
)
