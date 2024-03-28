package User

import (
	"github.com/Thenecromance/OurStories/base/SQL"
	"sync"
)

// when user login or register, or something else need user data, it will be looked like this

var (
	bindToTable sync.Once
)

type Info struct {
	ID          int    `json:"id" db:"id"`
	UserName    string `json:"username" db:"username, size:20" form:"username" binding:"required"`
	Password    string `json:"password" db:"password, size:64" form:"password" binding:"required"`
	Email       string `json:"email" db:"email, size:64" form:"email"`
	Gender      string `json:"gender" db:"gender"`
	CreatedAt   int    `json:"created_at" db:"created_at"`       // create time stamp
	LastLoginAt int    `json:"last_login_at" db:"last_login_at"` // last login time stamp
	Mate        int    `json:"mate" db:"mate"`                   // after setting the mate, it will be recorded here

}

func (i *Info) TableName() string {
	return "users"
}

func (i *Info) initdb() {

	// using sync.Once to make sure the table is only bind once
	bindToTable.Do(func() {
		tm := SQL.Default().AddTableWithName(Info{}, "users")
		tm.SetKeys(true, "ID")
		tm.ColMap("username").SetUnique(true).SetNotNull(true)
		tm.ColMap("email").SetUnique(true).SetNotNull(true)
		tm.ColMap("password").SetNotNull(true)
	})
	return
}

// SelfInsert will insert the self data into the database
func (i *Info) SelfInsert() error {
	i.initdb()
	return SQL.Default().Insert(i)
}

// SelfUpdate will update the self data into the database
func (i *Info) SelfUpdate() error {
	i.initdb()
	_, err := SQL.Default().Update(i)
	return err
}

// SelfDelete will delete the self data from the database
func (i *Info) SelfDelete() error {
	i.initdb()
	_, err := SQL.Default().Delete(i)
	return err
}

// SelfGet will get the self data from the database
func (i *Info) SelfGet() error {
	i.initdb()
	return SQL.Default().SelectOne(i, "SELECT * FROM `users` WHERE id = ? OR  email = ?", i.ID, i.Email)
}

// GetUserFromDatabase this method must be used in a empty object, it will override the object's data
func (i *Info) GetUserFromDatabase(userId int) {
	i.ID = userId
	i.SelfGet()
}
