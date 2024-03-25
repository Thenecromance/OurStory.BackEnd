package User

import "time"

type loginInfo struct {
	ID     int       `json:"id" db:"id"`
	UserId int       `json:"user_id" db:"user_id"`
	Date   time.Time `json:"date" db:"date"`
}
