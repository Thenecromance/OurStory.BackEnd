package models

import (
	"time"

	"github.com/Thenecromance/OurStories/utility/id"
	"github.com/Thenecromance/OurStories/utility/log"
	"gopkg.in/gorp.v2"
)

// =========================================================
// hooks
// =========================================================

func (i *Item) PreInsert(s gorp.SqlExecutor) error {
	i.ItemId = id.Generate()
	i.CreateAt = time.Now().UnixMilli()
	return nil
}

func (ub *UserBalance) PreInsert(s gorp.SqlExecutor) error {
	ub.UpdateAt = time.Now().UnixMilli()

	return nil
}

func (ub *UserBalance) PostInsert(s gorp.SqlExecutor) error {

	return s.Insert(&Transaction{
		UserId: ub.UserId,
		Amount: 100.0,
		Type:   "credit",
		Status: "completed",
	})
}

func (t *Transaction) PreInsert(s gorp.SqlExecutor) error {
	t.TransactionId = id.Generate()
	t.TimeStamp = time.Now().UnixMilli()

	return nil

}

func applyTransactionToUserBalance(s gorp.SqlExecutor, t *Transaction) error {
	balance := UserBalance{}
	if err := s.SelectOne(&balance, "SELECT * FROM UserBalances WHERE user_id = ?", t.UserId); err != nil {
		log.Warnf("Error getting balance: %v", err)
		return err
	}
	_, err := s.Update(&UserBalance{UserId: t.UserId, Balance: balance.Balance + t.Amount})
	if err != nil {
		log.Warnf("Error updating balance: %v", err)
		return err
	}

	return s.Insert(&TransactionLog{
		TransactionId: t.TransactionId,
		Message:       "Transaction created",
		TimeStamp:     t.TimeStamp,
	})
}

func (t *Transaction) PostInsert(s gorp.SqlExecutor) error {
	if err := applyTransactionToUserBalance(s, t); err != nil {
		return err
	}
	return nil
}

func (t *TransactionLog) PreInsert(s gorp.SqlExecutor) error {
	t.LogId = id.Generate()
	return nil
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
