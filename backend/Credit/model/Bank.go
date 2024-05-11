package model

import (
	"fmt"
	"github.com/Thenecromance/OurStories/backend/Credit/data"
	"github.com/Thenecromance/OurStories/base/SQL"
	"github.com/Thenecromance/OurStories/base/log"
	"gopkg.in/gorp.v2"
	"time"
)

type Bank struct {
	handler *gorp.DbMap
}

func (b *Bank) init() {
	log.Info("start to init Bank model")
	b.handler = SQL.Default()
	balance := data.Balance{}
	history := data.BalanceHistory{}
	balance.SetupTable(b.handler)
	history.SetupTable(b.handler)

	log.Info("init Bank model success")
}

// when user first time to use this bank, the bank will create an account for the user
func (b *Bank) CreateAccount(userId int) error {
	account := data.Balance{
		UserID:    userId,
		CreatedAt: time.Now().Unix(),
	}

	err := b.handler.Insert(&account)
	return err
}

func (b *Bank) DeleteAccount(userId int) {

}

func (b *Bank) GetBalance(userId int) *data.Balance {
	accounts, err := b.handler.Get(data.Balance{}, userId)
	if err != nil {
		log.Error("failed to get balance with error: %s", err.Error())
		return nil
	}
	return accounts.(*data.Balance)
}

func (b *Bank) Deposit(userId int, amount int) error {
	account := b.GetBalance(userId)
	if account == nil {
		log.Error("failed to get balance")
		return fmt.Errorf("failed to get balance")
	}

	return nil
}
