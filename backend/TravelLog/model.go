package TravelLog

import "time"

type Data struct {
	Id        int       `json:"id" db:"id"`
	Location  string    `json:"location" db:"location"`
	Message   string    `json:"message" db:"message"`
	ImagePath string    `json:"img" db:"img"`
	Date      time.Time `json:"date" db:"date"`
}

type Info struct {
}
