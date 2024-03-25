package User

import "time"

type userLoggedInInfo struct {
	Id     int       `json:"id" db:"id"`
	UserId int       `json:"user_id" db:"user_id"`
	Date   time.Time `json:"date" db:"date"`
}
