package User

import (
	"github.com/Thenecromance/OurStories/base/SQL"
	"github.com/Thenecromance/OurStories/base/hash"
	"sync"
	"sync/atomic"
	"time"
)

var (
	bindInfoTable sync.Once
)

type Info struct {
	encrypted   atomic.Bool `db:"-" json:"-"`
	Id          int         `db:"id" json:"id"`
	Avatar      string      `db:"avatar" json:"avatar"` // the path of avatar
	UserName    string      `db:"username,notnull" json:"username"`
	Password    string      `db:"password" json:"password"`
	Email       string      `db:"email"    json:"email"`
	MBTI        string      `db:"mbti" json:"mbti"`
	CreatedTime int64       `db:"created_time" json:"created_time"`
	LastLogin   int64       `db:"last_login" json:"last_login"`
	Gender      int         `db:"gender"`
}

func BindInfoTable() {
	bindInfoTable.Do(func() {
		tbl := SQL.Default().AddTableWithName(Info{}, "user")
		tbl.SetKeys(true, "Id")
		SQL.Default().CreateTablesIfNotExists()
	})

}

// Encrypt the Sensitive data in here, like password (but I still need consider about how to let each object only encrypt once)
func (i *Info) Encrypt() {
	if !i.encrypted.Load() {
		i.Password = hash.Hash(i.Password)
		i.encrypted.Store(true)
	}
}

func (i *Info) InsertToSQL() error {
	now := time.Now().Unix()
	i.CreatedTime = now

	return SQL.Default().Insert(i)
}

func (i *Info) UpdateToSQL() error {
	_, err := SQL.Default().Update(i)
	return err
}

func (i *Info) GetFromSQLByUserName() error {
	err := SQL.Default().SelectOne(i, "SELECT * FROM user WHERE username=?", i.UserName)
	if err != nil {
		return err
	}

	// when get from sql, the password is encrypted
	if i.Password != "" {
		i.encrypted.Store(true)
	}

	return nil
}
func (i *Info) GetFromSQLById() error {
	err := SQL.Default().SelectOne(i, "SELECT * FROM user WHERE id=?", i.Id)
	if err != nil {
		return err
	}

	// when get from sql, the password is encrypted
	if i.Password != "" {
		i.encrypted.Store(true)
	}

	return nil
}

func (i *Info) Copy() Info {
	return Info{
		Id:          i.Id,
		Avatar:      i.Avatar,
		UserName:    i.UserName,
		Password:    i.Password,
		Email:       i.Email,
		CreatedTime: i.CreatedTime,
		LastLogin:   i.LastLogin,
	}
}

func (i *Info) Overwrite(new Info) {
	if new.Avatar != "" {
		i.Avatar = new.Avatar
	}
	if new.UserName != "" {
		i.UserName = new.UserName
	}
	//if new.Password != "" {
	//	i.Password = new.Password
	//}
	if new.Email != "" {
		i.Email = new.Email
	}
	if new.MBTI != "" {
		i.MBTI = new.MBTI
	}

	if new.Gender != i.Gender {
		i.Gender = new.Gender
	}
}
