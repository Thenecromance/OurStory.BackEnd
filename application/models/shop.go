package models

import (
	"github.com/Thenecromance/OurStories/utility/id"
	"gopkg.in/gorp.v2"
	"time"
)

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

func (i *Item) PreInsert(s gorp.SqlExecutor) error {
	i.ItemId = id.Generate()
	i.CreateAt = time.Now().UnixMilli()
	return nil
}

type UserBalance struct {
	UserId   int     `json:"user_id,omitempty" db:"user_id"`
	Balance  float64 `json:"balance,omitempty" db:"balance"`
	UpdateAt int64   `json:"update_at,omitempty" db:"update_at"`
}

func (ub *UserBalance) PreInsert(s gorp.SqlExecutor) error {
	ub.UpdateAt = time.Now().UnixMilli()
	return nil
}

//=========================================================
// Transactions model
//=========================================================

type Transactions struct {
	TransactionId int64   `json:"transaction_id,omitempty" db:"transaction_id"`
	UserId        int64   `json:"user_id,omitempty" db:"user_id"`
	Amount        float64 `json:"item_id,omitempty" db:"item_id"`
	Type          string  `json:"type,omitempty" db:"transaction_type"`
	Status        string  `json:"status,omitempty" db:"status"`
	TimeStamp     int64   `json:"created_at,omitempty" db:"created_at"`
}

func (t *Transactions) PreInsert(s gorp.SqlExecutor) error {
	t.TransactionId = id.Generate()
	t.TimeStamp = time.Now().UnixMilli()

	return s.Insert(&TransactionLog{
		TransactionId: t.TransactionId,
		Message:       "Transaction created",
		TimeStamp:     t.TimeStamp,
	})

}

type TransactionLog struct {
	LogId         int64  `json:"log_id,omitempty" db:"log_id"`
	TransactionId int64  `json:"transaction_id,omitempty" db:"transaction_id"`
	Message       string `json:"message,omitempty" db:"log_message"`
	TimeStamp     int64  `json:"logged_at,omitempty" db:"logged_at"`
}

func (t *TransactionLog) PreInsert(s gorp.SqlExecutor) error {
	t.LogId = id.Generate()
	return nil
}

//=========================================================
// Carts model
//=========================================================

type Carts struct {
	CartId    int64 `json:"cart_id,omitempty" db:"cart_id"`       // Carts id
	UserId    int64 `json:"user_id,omitempty" db:"user_id"`       // user id
	CreatedAt int64 `json:"created_at,omitempty" db:"created_at"` // created at
}

func (c *Carts) PreInsert(s gorp.SqlExecutor) error {
	c.CartId = id.Generate()
	c.CreatedAt = time.Now().UnixMilli()
	return nil
}

func (c *Carts) PreDelete(s gorp.SqlExecutor) error {
	// delete all items in user's cart
	_, err := s.Exec("DELETE FROM CartedItems WHERE cart_id = ?", c.CartId)
	return err
}

type CartedItem struct {
	ItemId int64 `json:"item_id,omitempty" db:"item_id"` // item id
	Amount int   `json:"amount,omitempty" db:"amount"`   // amount of item
}

// =========================================================
// hooks
// =========================================================
