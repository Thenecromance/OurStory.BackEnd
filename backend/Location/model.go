package Location

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Thenecromance/OurStories/base/logger"
)

const (
	ipLocationApi = `https://restapi.amap.com/v3/ip?parameters&key=%s&ip=%s` //for using ip to get location
)

type Data struct {
	Status    string `json:"status"`
	Info      string `json:"info"`
	Infocode  string `json:"infocode"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Adcode    string `json:"adcode"`
	Rectangle string `json:"rectangle"`
	Ip        string `json:"ip"`
}

type Model struct {
	//db *gorp.DbMap

	locationMap map[string]Data //cache
}

func (m *Model) requestFromAMap(ip string) (result Data) {
	if m.locationMap == nil {
		m.locationMap = make(map[string]Data)
	}
	//prevent request the same ip
	if loc, ok := m.locationMap[ip]; ok {
		return loc
	}

	response, err := http.Get(fmt.Sprintf(ipLocationApi /* AMapToken.Instance().Amap */, "", ip))
	if err != nil {
		logger.Get().Errorf("Ip location Request failed,%s   .... %s", ip, err.Error())
		return
	}

	if response.StatusCode != 200 {
		logger.Get().Errorf("Ip location Response failed for %s. Response Code: %d", ip, response.StatusCode)
		return
	}
	defer response.Body.Close()

	buffer, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Get().Errorf("fail to ReadAll Ip location .... %s", err.Error())
		return
	}

	loc := Data{}
	err = json.Unmarshal(buffer, &loc)
	if err != nil {
		logger.Get().Errorf("fail to Unmarshal Ip location .... %s", err.Error())
		return
	}

	loc.Ip = ip
	m.locationMap[ip] = loc

	return m.locationMap[ip]
}

func (m *Model) requestFromCache(ip string) (Data, bool) {
	if m.locationMap == nil {
		m.locationMap = make(map[string]Data)
	}
	if loc, ok := m.locationMap[ip]; ok {
		return loc, true
	}
	return Data{}, false
}

func (m *Model) GetLocation(ip string) Data {
	loc, ok := m.requestFromCache(ip)
	if ok {
		return loc
	}
	return m.requestFromAMap(ip)
}
