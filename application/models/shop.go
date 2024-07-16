package models

//=========================================================
// Shop model
//=========================================================

type Item struct {
	ItemId      int64   `json:"item_id,omitempty" db:"item_id"`
	Name        string  `json:"name,omitempty" db:"name"`
	Description string  `json:"description,omitempty" db:"description"`
	Limit       int     `json:"limit,omitempty" db:"limit"`
	Price       float64 `json:"price,omitempty" db:"price"`
	ReleaseDate int64   `json:"release_date,omitempty" db:"release_date"`
	ExpireDate  int64   `json:"expire_date,omitempty" db:"expire_date"`
	CreateAt    int64   `json:"create_at,omitempty" db:"create_at"`
	Publisher   int     `json:"publisher,omitempty" db:"publisher"`
}
type UserBalance struct {
	UserId   int64   `json:"user_id,omitempty" db:"user_id"`
	Balance  float64 `json:"balance,omitempty" db:"balance"`
	UpdateAt int64   `json:"updated_at,omitempty" db:"updated_at"`
}

//=========================================================
// Transactions model
//=========================================================

type Transactions struct {
	TransactionId int64   `json:"transaction_id,omitempty" db:"transaction_id"`
	UserId        int64   `json:"user_id,omitempty" db:"user_id"`
	Amount        float64 `json:"amount,omitempty" db:"amount"`
	Type          string  `json:"type,omitempty" db:"transaction_type"`
	Status        string  `json:"status,omitempty" db:"status"`
	TimeStamp     int64   `json:"created_at,omitempty" db:"created_at"`
}

type TransactionLog struct {
	LogId         int64  `json:"log_id,omitempty" db:"log_id"`
	TransactionId int64  `json:"transaction_id,omitempty" db:"transaction_id"`
	Message       string `json:"message,omitempty" db:"log_message"`
	TimeStamp     int64  `json:"logged_at,omitempty" db:"logged_at"`
}

//=========================================================
// Carts model
//=========================================================

type Carts struct {
	CartId    int64 `json:"cart_id,omitempty" db:"cart_id"`       // Carts id
	UserId    int64 `json:"user_id,omitempty" db:"user_id"`       // user id
	CreatedAt int64 `json:"created_at,omitempty" db:"created_at"` // created at
}
