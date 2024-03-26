package User

type loginInfo struct {
	ID        int   `json:"id" db:"id"`
	UserId    int   `json:"user_id" db:"user_id"`
	TimeStamp int64 `json:"date" db:"stamp"`
}
