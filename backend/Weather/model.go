package Weather

import (
	"encoding/json"
	"fmt"
	"github.com/Thenecromance/OurStories/backend/AMapToken"
	"github.com/Thenecromance/OurStories/base/logger"
	"gopkg.in/gorp.v2"
	"io"
	"net/http"
)

const (
	weatherApiTemplate = `https://restapi.amap.com/v3/weather/weatherInfo?key=%s&city=%s`
)

/*type Weather struct {
	Date           string `json:"date"`
	Week           string `json:"week"`
	Dayweather     string `json:"dayweather"`
	Nightweather   string `json:"nightweather"`
	Daytemp        string `json:"daytemp"`
	Nighttemp      string `json:"nighttemp"`
	Daywind        string `json:"daywind"`
	Nightwind      string `json:"nightwind"`
	Daypower       string `json:"daypower"`
	Nightpower     string `json:"nightpower"`
	DaytempFloat   string `json:"daytemp_float"`
	NighttempFloat string `json:"nighttemp_float"`
}
type CityWeather struct {
	City       string    `json:"city"`
	Adcode     string    `json:"adcode"`
	Province   string    `json:"province"`
	Reporttime string    `json:"reporttime"`
	Weathers   []Weather `json:"casts"`
}
type AMapWeather struct {
	Status      string        `json:"status"`
	Count       string        `json:"count"`
	Info        string        `json:"info"`
	Infocode    string        `json:"infocode"`
	CityWeather []CityWeather `json:"forecasts"`
}*/

type Weather struct {
	Province         string `json:"province"`
	City             string `json:"city"`
	Adcode           string `json:"adcode"`
	Weather          string `json:"weather"`
	Temperature      string `json:"temperature"`
	Winddirection    string `json:"winddirection"`
	Windpower        string `json:"windpower"`
	Humidity         string `json:"humidity"`
	Reporttime       string `json:"reporttime"`
	TemperatureFloat string `json:"temperature_float"`
	HumidityFloat    string `json:"humidity_float"`
}

// get current weather from AMap
type WeatherReponse struct {
	Status   string    `json:"status"`
	Count    string    `json:"count"`
	Info     string    `json:"info"`
	Infocode string    `json:"infocode"`
	Lives    []Weather `json:"lives"`
}

type Model struct {
	db *gorp.DbMap

	weather map[string]Weather //cache key:code value:weather
}

func (m *Model) UpdateFromAMap(code string) (result Weather) {
	if m.weather == nil {
		m.weather = make(map[string]Weather)
	}

	resp, err := http.Get(fmt.Sprintf(weatherApiTemplate, AMapToken.Instance().Amap, code))
	if err != nil {
		logger.Get().Error("fail to request weather info", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		logger.Get().Error("Weather request failed with code :", resp.StatusCode)
		return
	}

	buffer, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Get().Error("fail to read weather info", err)
		return
	}

	var wResposne WeatherReponse
	err = json.Unmarshal(buffer, &wResposne)
	if err != nil {
		logger.Get().Error("fail to unmarshal weather info", err)
		return
	}
	if wResposne.Status != "1" {
		logger.Get().Error("fail to get weather info", wResposne.Info)
		return
	}
	// record in to the weather map

	m.weather[code] = wResposne.Lives[0]

	return m.weather[code]
}

func (m *Model) GetWeatherByCode(code string) Weather {
	if m.weather == nil {
		m.weather = make(map[string]Weather)
	}
	//check weather in cache
	if weather, ok := m.weather[code]; ok {
		logger.Get().Info("get weather from cache")
		return weather
	}

	// request from amap
	weather := m.UpdateFromAMap(code)

	m.weather[code] = weather

	return m.weather[code]
}
