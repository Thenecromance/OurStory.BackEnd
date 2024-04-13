package amap

import (
	"fmt"
	"github.com/Thenecromance/OurStories/backend/amap/data"
	"io"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/Thenecromance/OurStories/base/lru"
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

func (a *Amap) request(url string) (buffer []byte) {
	log.Debug("start to request :%s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("failed to request url : %s \nerror: %s", url, err.Error())
		return
	}
	if resp == nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("target server don't want to answer you with status code:%d", resp.StatusCode)
		return
	}

	buffer, err = io.ReadAll(resp.Body)

	if err != nil {
		log.Error("failed read response :%s", err.Error())
		return
	}

	return
}

// so far GetWeatherByIp only support the check the weather by it's Ip address. if want to support use city's name to get the weather
func (a *Amap) GetWeatherByIp(address string) *data.Weather {
	// prevent invalid operation before the config setup
	if !a.allowToUse {
		log.Warnf("missing shit config, please setup it then you can use it ")
		return nil
	}

	// first get location's adcode , amap only support it's owned ad code
	loc := a.getLocationByIP(address)
	if loc == nil {
		return nil
	}
	// then just directly return the weather data
	return a.getWeather(loc.Adcode)
}

func New() *Amap {
	ptr := &Amap{
		adcodeCache:  lru.New(100),
		weatherCache: lru.New(100),
	}
	ptr.initConfig()

	return ptr
}

func TestCase() {
	demo := New()

	buildIp := func(slice *[]string, start string) {
		// "1.2.0."
		for i := 0; i <= 3; i++ {

			*slice = append(*slice, start+strconv.Itoa(i))
		}
	}
	ipList := make([]string, 0)

	buildIp(&ipList, "1.2.0.")
	buildIp(&ipList, "1.0.1.")
	buildIp(&ipList, "1.12.0.")
	buildIp(&ipList, "1.56.0.")
	buildIp(&ipList, "1.12.0.")

	for i := 0; i < 100; i++ {

		fmt.Println(demo.GetWeatherByIp(ipList[rand.Intn(len(ipList))]))
		fmt.Println()
	}

}
