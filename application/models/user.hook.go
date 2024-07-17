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

func createUserRegisterAnniversary(s gorp.SqlExecutor, user *User) error {
	return s.Insert(&Anniversary{
		UserId:      user.UserId,
		Name:        "Reigster day!",
		Description: "welcome to this system",
	})
}
