package creator

import (
	"fmt"
	"github.com/Thenecromance/OurStories/backend/User/Data"
	"github.com/Thenecromance/OurStories/base/SQL"
	"time"
)

type Creator struct {
}

func (c *Creator) checkUserNameOrEmailInDb(username, email string) bool {
	err := SQL.Default().SelectOne(Data.Info{}, "SELECT * FROM user WHERE username=? or email=?", username, email)
	return err != nil
}
func (c *Creator) NewUser(username, email, password string) error {
	if c.checkUserNameOrEmailInDb(username, email) {
		return fmt.Errorf("username or email already exists")
	}

	info := &Data.Info{
		AuthorizationInfo: Data.AuthorizationInfo{
			UserName: username,
			Password: password,
		},
		CommonInfo: Data.CommonInfo{
			Email:       email,
			NickName:    username,
			CreatedTime: time.Now().Unix(),
			LastLogin:   time.Now().Unix(),
		},
	}

	err := info.InsertToSQL()
	if err != nil {
		return fmt.Errorf("insert user failed: %s", err)
	}
	return nil
}

func New() *Creator {
	return &Creator{}
}
