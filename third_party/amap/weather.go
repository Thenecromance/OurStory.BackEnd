package amap

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Thenecromance/OurStories/base/logger"
	"github.com/Thenecromance/OurStories/third_party/amap/data"
)

const (
	weatherApi = "https://restapi.amap.com/v3/weather/weatherInfo?city=%s&key=%s"
)

func (a *Amap) getWeather(adcode string) (result *data.Weather) {
	result = &data.Weather{}
	requestUri := fmt.Sprintf(weatherApi, adcode, a.getToken())
	logger.Get().Debugf("start to request %s", requestUri)
	buffer := a.request(requestUri)
	if buffer == nil {
		logger.Get().Error("weather reuest failed")
		return
	}

	if err := json.Unmarshal(buffer, result); err != nil {
		logger.Get().Errorf("fail to parse response ,%s", err)
		return
	}

	logger.Get().Debugf("%s \n response : %s", requestUri, string(buffer))

	a.weatherCache.Add(adcode, result, time.Now().Add(1*time.Hour))

	return
}
