package models

type Anniversary struct {
	Id                  int64  `json:"anniversary_id"         db:"anniversary_id"` // the anniversary's id
	UserId              int64  `json:"user_id"   db:"user_id"`                     // the user who create this anniversary
	Date                int64  `json:"anniversary_date" db:"anniversary_date"`     // the time when the anniversary happened
	Name                string `json:"title"      db:"title"`                      // the name of the anniversary
	Description         string `json:"description"       db:"description"`         // the information of the anniversary
	CreatedTime         int64  `json:"created_at" db:"created_at"`                 // the time when the anniversary is created
	UpdateAt            int64  `json:"updated_at" db:"updated_at"`                 // the time when the anniversary is updated
	SharedWithMarshaled string `json:"-" db:"shared_with"`                         // the user list who will share this anniversary
	// these fields are not in the database which will be calculated by the server
	SharedWith []int `json:"shared_with" db:"-"`  // the user list who will share this anniversary
	TotalSpend int   `json:"total_spend" db:"-"`  // this filed will be calculated by the server until now
	TimeToNext int   `json:"time_to_next" db:"-"` // this filed will be calculated by the server until the next anniversary
}
