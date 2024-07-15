package models

//=========================================================
// Shop model
//=========================================================

type Item struct {
	ItemId      int     `json:"item_id,omitempty" db:"item_id"`
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
}

//=========================================================
// Transactions model
//=========================================================

type Transactions struct {
}
type TransactionLog struct {
}

//=========================================================
// Carts model
//=========================================================

type Carts struct {
	CartId    int64 `json:"cart_id,omitempty" db:"cart_id"`       // Carts id
	UserId    int64 `json:"user_id,omitempty" db:"user_id"`       // user id
	CreatedAt int64 `json:"created_at,omitempty" db:"created_at"` // created at
}

type CartedItem struct {
	ItemId int64 `json:"item_id,omitempty" db:"item_id"` // item id
	Amount int   `json:"amount,omitempty" db:"amount"`   // amount of item

}
