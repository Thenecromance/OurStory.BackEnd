package amap

import (
	"io"
	"net/http"

	"github.com/Thenecromance/OurStories/base/logger"
	"github.com/Thenecromance/OurStories/base/lru"
	"github.com/Thenecromance/OurStories/third_party/amap/data"
)

type Amap struct {
	Token  string `json:"token"`
	UseSig bool   `json:"signature"`

	allowToUse   bool       `json:"-"`
	adcodeCache  *lru.Cache `json:"-"`
	weatherCache *lru.Cache `json:"-"`
}

func (a *Amap) getToken() string {
	return a.Token
}

func (a *Amap) GetWeatherByIp(address string) *data.Weather {
	loc := a.getLocationByIP(address)

	return a.getWeather(loc.Adcode)
}

func (a *Amap) request(url string) (buffer []byte) {
	logger.Get().Debug("start to request :%s", url)
	resp, err := http.Get(url)
	if err != nil {
		logger.Get().Errorf("failed to request url : %s \nerror: %s", url, err.Error())
		return
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		logger.Get().Errorf("target server don't want to answer you with status code:%d", resp.StatusCode)
		return
	}

	buffer, err = io.ReadAll(resp.Body)

	if err != nil {
		logger.Get().Error("failed read response :%s", err.Error())
		return
	}

	return
}

func New() *Amap {
	ptr := &Amap{
		adcodeCache:  lru.New(0),
		weatherCache: lru.New(0),
	}

	return ptr
}
