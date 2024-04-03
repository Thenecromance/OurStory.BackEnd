package Validator

import (
	"github.com/Thenecromance/OurStories/backend/User/Data"
	"github.com/Thenecromance/OurStories/base/SQL"
)

type Validator struct {
}

func (v *Validator) checkUserNameOrEmailInDb(username string) bool {
	err := SQL.Default().SelectOne(Data.Info{}, "SELECT * FROM user WHERE username=? or email=?", username, username)
	return err != nil
}

func (v *Validator) ValidateUser(info *Data.AuthorizationInfo, dbInfo *Data.Info) bool {
	info.Encrypt()
	dbInfo.Encrypt()

	if info.UserName != dbInfo.UserName {
		return false
	}
	if info.Password != dbInfo.Password {
		return false
	}

	return true
}

func New() *Validator {
	return &Validator{}
}
