package models

import (
	"time"

	"github.com/Thenecromance/OurStories/utility/log"
	"gopkg.in/gorp.v2"
)

func (u *User) PreInsert(s gorp.SqlExecutor) error {
	// u.UserId = id.Generate()  // no need to generate id, because it is auto increment
	u.CreatedAt = time.Now().UnixMilli()
	u.LastLogin = time.Now().UnixMilli()

	return nil
}

func (u *User) PostInsert(s gorp.SqlExecutor) error {
	// get the user id
	s.SelectOne(&u.UserId, "SELECT user_id FROM Users WHERE username = ? and email = ?", u.UserName, u.Email)
	go func() {
		logUserRegister(s, u)
		s.Insert(&UserBalance{
			UserId:  u.UserId,
			Balance: 0.0,
		})
		err := createUserRegisterAnniversary(s, u)
		if err != nil {
			log.Error(err)
		}
	}()
	return nil
}

func (ll *LoginLogs) PreInsert(s gorp.SqlExecutor) error {
	ll.LoginTime = time.Now().UnixMilli()
	return nil
}
func (ll *LoginLogs) PreUpdate(s gorp.SqlExecutor) error {
	ll.LoginTime = time.Now().UnixMilli()
	return nil
}

func logUserRegister(s gorp.SqlExecutor, user *User) error {
	return s.Insert(&LoginLogs{
		UserId:    user.UserId,
		LoginTime: time.Now().UnixMilli(),
	})
}

/*
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
*/

func createUserRegisterAnniversary(s gorp.SqlExecutor, user *User) error {
	return s.Insert(&Anniversary{
		UserId:      user.UserId,
		Name:        "Reigster day!",
		Description: "welcome to this system",
	})
}
