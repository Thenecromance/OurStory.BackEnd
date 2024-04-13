package amap

import (
	"encoding/json"
	"fmt"
	"github.com/Thenecromance/OurStories/backend/amap/data"
	"time"
)

const (
	weatherApi = "https://restapi.amap.com/v3/weather/weatherInfo?city=%s&key=%s"
)

func (a *Amap) getWeather(adcode string) (result *data.Weather) {
	// prevent invalid operation before the config setup
	if !a.allowToUse {
		log.Warnf("missing shit config, please setup it then you can use it ")
		return
	}

	cache := a.getWeatherFromCache(adcode)
	if cache != nil {
		return cache

	}

	requestUri := fmt.Sprintf(weatherApi, adcode, a.getToken())
	log.Debugf("start to request %s", requestUri)
	buffer := a.request(requestUri)
	if buffer == nil {
		log.Error("weather reuest failed")
		return
	}

	resp := &data.WeatherReponse{}
	if err := json.Unmarshal(buffer, resp); err != nil {
		log.Errorf("fail to parse response ,%s", err)
		return
	}

	log.Debugf("%s \n response : %s", requestUri, string(buffer))

	if resp.Status != "1" {
		log.Error("fail to get the result ")
		return
	}

	// result = resp.Lives[0]

	result = resp.Lives[0].Copy()
	a.weatherCache.Add(adcode, resp.Lives[0].Copy(), time.Now().Add(1*time.Hour))

	return
}

func (a *Amap) getWeatherFromCache(adcode string) (result *data.Weather) {
	ptr, exist := a.weatherCache.Get(adcode)

	if !exist {
		return
	}
	result = ptr.(*data.Weather).Copy()
	return

}
