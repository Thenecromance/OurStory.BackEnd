package Data

type CommonInfo struct {
	Id          int    `db:"id"                    json:"id"`
	Avatar      string `db:"avatar"                json:"avatar"            form:"avatar"`  // the path of avatar
	NickName    string `db:"nickname"              json:"nickname"         form:"nickname"` // nickname is the name that show to others
	Email       string `db:"email"                 json:"email"            form:"email"`
	MBTI        string `db:"mbti"                  json:"mbti"             form:"mbti"`
	Birthday    int64  `db:"birthday"              json:"birthday"         form:"birthday"`
	CreatedTime int64  `db:"created_time"          json:"created_time"`
	LastLogin   int64  `db:"last_login"            json:"last_login"`
	Gender      int    `db:"gender"                json:"gender"            form:"gender"`
}

func (ci *CommonInfo) ApplyNewInfo(info *CommonInfo) {
	if info.Avatar != "" {
		ci.Avatar = info.Avatar
	}
	if info.NickName != "" {
		ci.NickName = info.NickName
	}
	if info.Email != "" {
		ci.Email = info.Email
	}
	if info.MBTI != "" {
		ci.MBTI = info.MBTI
	}
	if info.Birthday != 0 {
		ci.Birthday = info.Birthday
	}
	if info.Gender != 0 {
		ci.Gender = info.Gender
	}
}
