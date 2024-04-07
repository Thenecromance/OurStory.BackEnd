package data

type LocationResponse struct {
	Status string `json:"status"`
	Info   string `json:"info"`
	// Rectangle string `json:"rectangle"`

	Location
}

type Location struct {
	Infocode string `json:"infocode"`
	Province string `json:"province"`
	City     string `json:"city"`
	Adcode   string `json:"adcode"`
}
