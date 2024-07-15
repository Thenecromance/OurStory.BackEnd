package models

type AnniversaryInDb struct {
	Id          int    `json:"anniversary_id"         db:"anniversary_id"` // the anniversary's id
	UserId      int    `json:"user_id"   db:"user_id"`                     // the user who create this anniversary
	Date        int64  `json:"anniversary_date" db:"anniversary_date"`     // the time when the anniversary happened
	Name        string `json:"title"      db:"title"`                      // the name of the anniversary
	Description string `json:"description"       db:"description"`         // the information of the anniversary
	CreatedTime int64  `json:"created_time" db:"created_time"`             // the time when the anniversary is created
	UpdateAt    int64  `json:"update_at" db:"update_at"`                   // the time when the anniversary is updated
	SharedWith  string `json:"shared_with" db:"shared_with"`               // the user list who will share this anniversary
}

type Anniversary struct {
	Id          int    `json:"id"         db:"id"`             // the anniversary's id
	UserId      int    `json:"user_id"   db:"user_id"`         // the user who create this anniversary
	TimeStamp   int64  `json:"time_stamp" db:"time_stamp"`     // the time when the anniversary happened
	Name        string `json:"name"      db:"name"`            // the name of the anniversary
	Info        string `json:"info"       db:"info"`           // the information of the anniversary
	CreatedTime int64  `json:"created_time" db:"created_time"` // the time when the anniversary is created
	SharedWith  []int  `json:"shared_with" db:"shared_with"`   // the user list who will share this anniversary
}

type AnniversaryDTO struct {
	AnniversaryInDb
	TotalSpend int `json:"total_spend"`  // this filed will be calculated by the server until now
	TimeToNext int `json:"time_to_next"` // this filed will be calculated by the server until the next anniversary
}
