package amap

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Thenecromance/OurStories/base/logger"
	"github.com/Thenecromance/OurStories/third_party/amap/data"
)

const (
	locationApi = "https://restapi.amap.com/v3/ip?ip=%s&key=%s"
)

func (a *Amap) getLocationByIP(addr string) (result *data.Location) {
	var exist bool
	result, exist = a.locationFromCache(addr)
	if exist {
		return result
	}

	requestUri := fmt.Sprintf(locationApi, addr, a.getToken())
	logger.Get().Debugf("location request url: %s", requestUri)

	buffer := a.request(requestUri)
	if buffer == nil {
		return
	}

	// result = &data.Location{}
	resp := &data.LocationResponse{}

	if json.Unmarshal(buffer, resp) != nil {
		logger.Get().Error("fail to parse the data")
		return
	}

	if resp.Status != "1" {
		logger.Get().Errorf("fail to request location , resp.Code = %s , info = %s", resp.Status, resp.Info)
		return
	}

	result = &data.Location{
		Infocode: resp.Infocode,
		Province: resp.Province,
		City:     resp.City,
		Adcode:   resp.Adcode,
	}

	a.adcodeCache.Add(addr, result, time.Now().Add(1*time.Hour))
	logger.Get().Debug("request location complete!")
	return
}

func (a *Amap) locationFromCache(addr string) (result *data.Location, exist bool) {

	logger.Get().Debugf("trying to get %s's adcode from cache ", addr)
	var ptr any
	ptr, exist = a.adcodeCache.Get(addr)
	if !exist {
		logger.Get().Debug("%s's info does not cached....")
		return
	}

	result = ptr.(*data.Location)
	logger.Get().Debugf("get  %s's cache complete! %s", result)
	return
}
