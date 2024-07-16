package models

import (
	"github.com/Thenecromance/OurStories/utility/id"
	"github.com/Thenecromance/OurStories/utility/log"
	"gopkg.in/gorp.v2"
	"time"
)

func (u *User) PreInsert(s gorp.SqlExecutor) error {
	u.UserId = id.Generate()
	u.CreatedAt = time.Now().UnixMilli()
	u.LastLogin = time.Now().UnixMilli()

	return nil
}

func (u *User) PostInsert(s gorp.SqlExecutor) error {
	// Create Balance account
	err := s.Insert(&UserBalance{
		UserId:  u.UserId,
		Balance: 0,
	})
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.Insert(&LoginLogs{
		UserId: u.UserId,
	})
	if err != nil {
		log.Error(err)
		return err
	}
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
