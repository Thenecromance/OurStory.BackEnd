package Data

import (
	"github.com/Thenecromance/OurStories/base/hash"
	"sync/atomic"
)

type AuthorizationInfo struct {
	encrypted atomic.Bool `db:"-"                   json:"-"`
	Password  string      `db:"password"            json:"password"         form:"password"`
	UserName  string      `db:"username,notnull"    json:"username"         form:"username"` // username is the name that use to login
}

// Encrypt the Sensitive data in here, like password (but I still need consider about how to let each object only encrypt once)
func (i *AuthorizationInfo) Encrypt() {
	if !i.encrypted.Load() {
		i.Password = hash.Hash(i.Password)
		i.encrypted.Store(true)
	}
}
