package amap

import (
	"encoding/json"
	"fmt"
	"github.com/Thenecromance/OurStories/backend/amap/data"
	"time"
)

const (
	locationApi = "https://restapi.amap.com/v3/ip?ip=%s&key=%s"
)

func (a *Amap) getLocationByIP(addr string) (result *data.Location) {
	// prevent invalid operation before the config setup
	if !a.allowToUse {
		log.Warnf("missing shit config, please setup it then you can use it ")
		return
	}

	//trying to read the ip info from the cache(once the user requested, it will cached in memory)
	var exist bool
	result, exist = a.locationFromCache(addr)
	// if found the info , no necessary to request from amap services
	if exist {
		return result
	}

	// build the request url
	requestUri := fmt.Sprintf(locationApi, addr, a.getToken())
	log.Debugf(requestUri)
	// do a request
	buffer := a.request(requestUri)
	if buffer == nil {
		return
	}

	// result = &data.Location{}

	loc := &data.Location{}
	resp := &data.LocationResponse{}
	resp.Location = loc
	if json.Unmarshal(buffer, resp) != nil {
		log.Error("fail to parse the data")
		return
	}

	if resp.Status != "1" {
		log.Errorf("fail to request location , resp.Code = %s , info = %s", resp.Status, resp.Info)
		return
	}

	// after all shit reuqest finished just pushed the result into the cache
	a.adcodeCache.Add(addr, loc, time.Now().Add(1*time.Hour))
	log.Debug("request location complete!")
	result = loc
	return
}

func (a *Amap) locationFromCache(addr string) (result *data.Location, exist bool) {

	log.Debugf("trying to get %s's adcode from cache ", addr)
	var ptr any
	ptr, exist = a.adcodeCache.Get(addr)
	if !exist {
		log.Debugf("%s's info does not cached....", addr)
		return
	}

	result = ptr.(*data.Location)
	log.Debugf("get  %s's cache complete! %s", addr, *result)
	return
}
