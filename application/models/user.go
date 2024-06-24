package models

// UserClaim only for signature JWT Token or other token that need to be signed
type UserClaim struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
}

// UserBasicDTO is the basic information of user
type UserBasicDTO struct {
	Id       int    `json:"id"               db:"id"`
	UserName string `json:"username"         db:"username"` // username is the name that use to login
	Avatar   string `json:"avatar"           db:"avatar"`   // the path of avatar
	NickName string `json:"nickname"         db:"nickname"` // nickname is the name that show to others
}

// UserAdvancedDTO is the advanced information of user
type UserAdvancedDTO struct {
	UserBasicDTO
	Email       string `db:"email"                 json:"email"             `
	MBTI        string `db:"mbti"                  json:"mbti"              `
	Birthday    int64  `db:"birthday"              json:"birthday"          `
	Gender      int    `db:"gender"                json:"gender"            `
	CreatedTime int64  `db:"created_time"          json:"created_time"      `
	LastLogin   int64  `db:"last_login"            json:"last_login"        `
}

// User is full user information
type User struct {
	UserAdvancedDTO
	Password string `db:"password"              json:"password"          `
	Salt     string `db:"salt"                  json:"salt"              `
}

// when user login, they need to provide username and password
type UserLogin struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UserRegister struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email"    form:"email"   `
}
