package User

import (
	"github.com/Thenecromance/OurStories/base/SQL"
	"sync"
)

type loginInfo struct {
	ID        int   `json:"id" db:"id"`
	UserId    int   `json:"user_id" db:"user_id"`
	TimeStamp int64 `json:"date" db:"stamp"`
}

var (
	bindLoginInfo sync.Once
)

func (i *loginInfo) TableName() string {
	return "loginInfo"
}

func (i *loginInfo) initdb() {

	// using sync.Once to make sure the table is only bind once
	bindToTable.Do(func() {
		tm := SQL.Default().AddTableWithName(Info{}, i.TableName())
		tm.SetKeys(true, "ID")
		tm.ColMap("username").SetUnique(true).SetNotNull(true)
		tm.ColMap("email").SetUnique(true).SetNotNull(true)
		tm.ColMap("password").SetNotNull(true)
	})
	return
}

// SelfInsert will insert the self data into the database
func (i *loginInfo) SelfInsert() error {
	i.initdb()
	return SQL.Default().Insert(i)
}

// SelfUpdate will update the self data into the database
func (i *loginInfo) SelfUpdate() error {
	i.initdb()
	_, err := SQL.Default().Update(i)
	return err
}

// SelfDelete will delete the self data from the database
func (i *loginInfo) SelfDelete() error {
	i.initdb()
	_, err := SQL.Default().Delete(i)
	return err
}

// SelfGet will get the self data from the database
func (i *loginInfo) SelfGet() error {
	i.initdb()
	return SQL.Default().SelectOne(i, "SELECT * FROM `loginInfo` WHERE id = ? OR  user_id = ?", i.ID, i.UserId)
}

// GetUserFromDatabase this method must be used in a empty object, it will override the object's data
func (i *loginInfo) GetUserFromDatabase(userId int) {
	i.ID = userId
	i.SelfGet()
}
