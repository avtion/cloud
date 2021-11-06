package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

var _ larkContentBuilder = (*weather)(nil)

func newWeather(cfg *weatherConfig) *weather {
	return &weather{cfg: cfg}
}

// weather 获取天气有关的内容
type weather struct {
	cfg *weatherConfig
}

// getLocalWeather 获取天气数据
func (w *weather) getLocalWeather() (*weatherResp, error) {
	const authScheme = "APPCODE"
	const api = "https://freecityid.market.alicloudapi.com/whapi/json/alicityweather/briefforecast3days"

	resp, err := resty.New().R().
		EnableTrace().
		SetAuthScheme(authScheme).
		SetAuthToken(w.cfg.AppCode).
		SetFormData(map[string]string{
			"cityId": w.cfg.CityID,
			"token":  "677282c2f1b3d718152c4e25ed434bc4",
		}).
		Post(api)
	if err != nil {
		return nil, err
	}
	if code := resp.StatusCode(); code != http.StatusOK {
		return nil, fmt.Errorf("ask weather failed, code: %d", code)
	}

	respWeather := new(weatherResp)
	if errUnmarshal := json.Unmarshal(resp.Body(), respWeather); errUnmarshal != nil {
		return nil, errUnmarshal
	}

	if respWeather.Code != 0 {
		return nil, errors.New(respWeather.Msg)
	}
	return respWeather, nil
}

// build implement larkContentBuilder
func (w *weather) build(_ *lark) []*larkElement {
	forecast, err := w.getLocalWeather()
	if err != nil || forecast.Data == nil || forecast.Data.Forecast[0] == nil {
		log.Printf("failed to get local weather info, err: %v, forcast: %#v", err, forecast)
		return nil
	}
	cityName := forecast.Data.City.Name
	todayForecast := forecast.Data.Forecast[0]
	content := fmt.Sprintf("☁️ **%s - %s** \n☀️ 白天温度：%s - %s \n🌛 夜间温度：%s - %s", cityName, todayForecast.ConditionDay, todayForecast.TempDay, todayForecast.WindDirDay, todayForecast.TempNight, todayForecast.WindDirNight)
	return []*larkElement{newLarkMarkdown(content)}
}
