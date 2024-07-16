package models

import (
	"time"

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
	go logUserRegister(s, u)
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
