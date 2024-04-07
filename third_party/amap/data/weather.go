package data

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

func (w Weather) Copy() *Weather {
	return &w
}

// get current weather from AMap
type WeatherReponse struct {
	Status   string    `json:"status"`
	Count    string    `json:"count"`
	Info     string    `json:"info"`
	Infocode string    `json:"infocode"`
	Lives    []Weather `json:"lives"`
}

//https://restapi.amap.com/v3/ip?ip=114.247.50.2&output=xml&key=4b25361ff579ef49eb044a818850a842
