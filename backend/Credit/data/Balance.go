package data

import (
	"github.com/Thenecromance/OurStories/base/log"
	"gopkg.in/gorp.v2"
)

type Balance struct {
	Id        int   `json:"id"       db:"id"`           // Balance ID
	UserID    int   `json:"user_id"  db:"user_id"`      // User ID
	Amount    int   `json:"amount"   db:"amount"`       // The amount of the balance
	CreatedAt int64 `json:"created_at" db:"created_at"` // The time when the balance was created
}

func (b *Balance) SetupTable(m *gorp.DbMap) {
	log.Info("start to binding object with table balance")
	t := m.AddTableWithName(Balance{}, "balance")
	t.SetKeys(true, "Id")

	err := m.CreateTablesIfNotExists()
	if err != nil {
		log.Errorf("failed to create [Bank] table with error: %s", err.Error())
		return
	}
}

type BalanceHistory struct {
	Id        int    `json:"id" db:"id"`                 // Balance history ID
	UserID    int    `json:"user_id" db:"user_id"`       // User ID
	Amount    int    `json:"amount" db:"amount"`         // The amount of the balance if the amount is negative means the balance is decreased
	Reason    string `json:"reason" db:"reason"`         // The reason for the change
	CreatedAt string `json:"created_at" db:"created_at"` // The time when the balance was changed
}

func (b *BalanceHistory) SetupTable(m *gorp.DbMap) {
	log.Info("start to binding object with table balance_history")
	t := m.AddTableWithName(BalanceHistory{}, "balance_history")
	t.SetKeys(true, "Id")
	t.ColMap("UserID").SetNotNull(true)

	err := m.CreateTablesIfNotExists()
	if err != nil {
		log.Errorf("failed to create [BalanceHistory] table with error: %s", err.Error())
		return
	}
}
