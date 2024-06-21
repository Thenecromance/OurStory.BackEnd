package models

type Anniversary struct {
	Id        string `json:"id"         db:"id"`
	OwnerId   int    `json:"owner_id"   db:"owner_id"`
	Owner     string `json:"owner"      db:"owner"`
	TimeStamp int64  `json:"time_stamp" db:"time_stamp"`
	Title     string `json:"title"      db:"title"`
	Info      string `json:"info"       db:"info"`
}

type AnniversaryDTO struct {
	Anniversary
	TotalSpend int `json:"total_spend"`  // this filed will be calculated by the server until now
	TimeToNext int `json:"time_to_next"` // this filed will be calculated by the server until the next anniversary
}
