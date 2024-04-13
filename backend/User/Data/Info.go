package Data

import (
	"github.com/Thenecromance/OurStories/base/SQL"
	"github.com/Thenecromance/OurStories/base/hash"
	"github.com/Thenecromance/OurStories/base/log"
	"sync"
	"time"
)

var (
	bindInfoTable sync.Once
)

type Info struct {
	AuthorizationInfo
	CommonInfo
}

func BindInfoTable() {
	bindInfoTable.Do(func() {
		tbl := SQL.Default().AddTableWithName(Info{}, "user")
		tbl.SetKeys(true, "Id")
		err := SQL.Default().CreateTablesIfNotExists()
		if err != nil {
			log.Error("Create table user failed: ", err)
			return
		}
	})

}

func (i *Info) Invalid() bool {
	return i.Id == 0
}

func (i *Info) InsertToSQL() error {
	now := time.Now().Unix()
	i.CreatedTime = now
	i.LastLogin = now
	i.Birthday = now

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
		AuthorizationInfo: AuthorizationInfo{
			UserName: i.UserName,
			Password: i.Password,
		},

		CommonInfo: CommonInfo{
			Id:          i.Id,
			Avatar:      i.Avatar,
			NickName:    i.NickName,
			Email:       i.Email,
			MBTI:        i.MBTI,
			CreatedTime: i.CreatedTime,
			LastLogin:   i.LastLogin,
		},
	}
}

func (i *Info) Overwrite(new *Info) {
	if new.Avatar != "" {
		i.Avatar = new.Avatar
	}
	if new.UserName != "" {
		i.UserName = new.UserName
	}
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

func (i *Info) ComparePassword(info *AuthorizationInfo) bool {
	return i.Password == hash.Hash(info.Password)
}

func (i *Info) GetFromSQLByUserNameOrEmail() error {
	err := SQL.Default().SelectOne(i, "SELECT * FROM user WHERE username=? OR email=?", i.UserName, i.Email)
	if err != nil {
		return err
	}

	if i.Password != "" {
		i.encrypted.Store(true)
	}
	return nil
}
