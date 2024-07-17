package models

const (
	RoleUser = iota
	RoleAdmin
	RoleMaster
)

// UserClaim only for signature JWT Token or other token that need to be signed
type UserClaim struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
}

// UserBasicDTO is the basic information of user
type UserBasicDTO struct {
	UserId   int64  `json:"user_id"               db:"user_id"`
	UserName string `json:"username"         db:"username"` // username is the name that use to login
	Avatar   string `json:"avatar"           db:"avatar"`   // the path of avatar
	NickName string `json:"nickname"         db:"nickname"` // nickname is the name that show to others
	Role     int    `json:"role"             db:"role"`
}

// UserAdvancedDTO is the advanced information of user
type UserAdvancedDTO struct {
	UserBasicDTO
	Email     string `db:"email"                 json:"email"             `
	Birthday  int64  `db:"birthday"              json:"birthday"          `
	Gender    string `db:"gender"                json:"gender"            `
	CreatedAt int64  `db:"created_at"          json:"created_at"      `
	LastLogin int64  `db:"-"            json:"last_login"        `
}

// User is full user information
type User struct {
	UserAdvancedDTO
	Password string `db:"pass_word"              json:"pass_word"          `
	Salt     string `db:"salt"                  json:"salt"              `
}

// hooks

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

type LoginLogs struct {
	UserId    int64 `json:"user_id" db:"user_id"`
	LoginTime int64 `json:"login_time" db:"login_time"`
	//LoginIp   string `json:"login_ip" db:"login_ip"`
}
