package user

import (
	"sync"
)

var (
	bindInfoTable sync.Once
)

type UserResponse struct {
	Id       int    `json:"id"`
	UserName string `json:"username"         ` // username is the name that use to login
	Avatar   string `json:"avatar"           ` // the path of avatar
	NickName string `json:"nickname"         ` // nickname is the name that show to others

}
type UserClaim struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
}

type UserInDb struct {
	Id          int    `db:"id"                    json:"id"`
	Password    string `db:"password"              json:"password"         `
	UserName    string `db:"username,notnull"      json:"username"         ` // username is the name that use to login
	Avatar      string `db:"avatar"                json:"avatar"           ` // the path of avatar
	NickName    string `db:"nickname"              json:"nickname"         ` // nickname is the name that show to others
	Email       string `db:"email"                 json:"email"            `
	MBTI        string `db:"mbti"                  json:"mbti"             `
	Birthday    int64  `db:"birthday"              json:"birthday"         `
	CreatedTime int64  `db:"created_time"          json:"created_time"`
	LastLogin   int64  `db:"last_login"            json:"last_login"`
	Gender      int    `db:"gender"                json:"gender"           `
}
