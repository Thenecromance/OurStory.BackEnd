package Bank

import (
	"encoding/json"
	"github.com/Thenecromance/OurStories/base/SQL"
	"github.com/Thenecromance/OurStories/base/logger"
	"gopkg.in/gorp.v2"
	"os"
)

type Account struct {
	Id     int `json:"id" db:"id"`
	UserID int `json:"user_id" db:"user_id"` //account's owner
	Credit int `json:"credit" db:"credit"`   //account's credit

}

// binding the table with gorp
func (d Account) setUpTable(db *gorp.DbMap) error {
	logger.Get().Info("start to binding Data with table bank")
	tbl := db.AddTableWithName(d, "bank")
	tbl.SetKeys(true, "Id")
	tbl.ColMap("UserID").SetNotNull(true)
	tbl.ColMap("Credit").SetNotNull(true)
	return db.CreateTablesIfNotExists()
}

type Model struct {
	newAccountSetting Account
	banks             []Account
}

func defaultAccount() Account {
	return Account{
		UserID: 0,
		Credit: 0,
	}
}
func (m *Model) loadDefaultAccountSetting() {
	bytes, err := os.ReadFile("defaultAccount.json")
	if err != nil {
		logger.Get().Error("load default account setting failed")
		return
	}
	err = json.Unmarshal(bytes, &m.newAccountSetting)
	if err != nil {
		logger.Get().Error("load default account setting failed")
		return
	}

}
func (m *Model) init() {
	Account{}.setUpTable(SQL.Default())
	m.loadDefaultAccountSetting()
}

// create a new account for a user
func (m *Model) createNewAccount(userID int) {

}

func (m *Model) checkUserCredit(userID int) {

}

func (m *Model) creditChanged(userID int, credit int) {

}
