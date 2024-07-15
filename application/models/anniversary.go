package models

import (
	"encoding/json"
	"github.com/Thenecromance/OurStories/utility/id"
	"gopkg.in/gorp.v2"
	"time"
)

type Anniversary struct {
	Id                  int64  `json:"anniversary_id"         db:"anniversary_id"` // the anniversary's id
	UserId              int64  `json:"user_id"   db:"user_id"`                     // the user who create this anniversary
	Date                int64  `json:"anniversary_date" db:"anniversary_date"`     // the time when the anniversary happened
	Name                string `json:"title"      db:"title"`                      // the name of the anniversary
	Description         string `json:"description"       db:"description"`         // the information of the anniversary
	CreatedTime         int64  `json:"created_time" db:"created_time"`             // the time when the anniversary is created
	UpdateAt            int64  `json:"update_at" db:"update_at"`                   // the time when the anniversary is updated
	SharedWithMarshaled string `json:"-" db:"shared_with"`                         // the user list who will share this anniversary
	// these fields are not in the database which will be calculated by the server
	SharedWith []int `json:"shared_with" db:"-"`  // the user list who will share this anniversary
	TotalSpend int   `json:"total_spend" db:"-"`  // this filed will be calculated by the server until now
	TimeToNext int   `json:"time_to_next" db:"-"` // this filed will be calculated by the server until the next anniversary
}

//------------------------------------------------------------
// hooks
//------------------------------------------------------------

func (a *Anniversary) PreInsert(s gorp.SqlExecutor) error {
	a.Id = id.Generate()
	a.CreatedTime = time.Now().UnixMilli()

	buf, err := json.Marshal(a.SharedWith)
	if err != nil {
		a.SharedWithMarshaled = ""
	}
	a.SharedWithMarshaled = string(buf)
	return nil
}

func (a *Anniversary) PostGet(s gorp.SqlExecutor) error {
	err := json.Unmarshal([]byte(a.SharedWithMarshaled), &a.SharedWith)
	return err
}
