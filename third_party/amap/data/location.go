package data

type LocationResponse struct {
	Status string `json:"status"`
	Info   string `json:"info"`
	// Rectangle string `json:"rectangle"`

	*Location // split the useful info from whole response
}

// due to a map's response cnotain lots of useless data(for most usage) so just split it to standard alone
type Location struct {
	Infocode string `json:"infocode"`
	Province string `json:"province"`
	City     string `json:"city"`
	Adcode   string `json:"adcode"`
}

func (loc Location) Copy() *Location {
	return &loc
}
