package Credit

import "gopkg.in/gorp.v2"

// when a user using credit to buy something, the credit will be modified
type Modified struct {
	Id       int `json:"id"       db:"id"                     `
	UserId   int `json:"user"     db:"user"                   binding:"required"` //for identifying the user who use the credit
	Amount   int `json:"amount"   db:"amount"     form:""     binding:"required"` //how much does the user use
	ReasonId int `json:"reason"   db:"reason"     form:""     binding:"required"` //the reason why the user use the credit
}

func (m Modified) setUpTable(db *gorp.DbMap) {
	tbl := db.AddTableWithName(m, "credit_modified")
	tbl.SetKeys(true, "Id")
	tbl.ColMap("UserId").SetNotNull(true)
	tbl.ColMap("Amount").SetNotNull(true)
	tbl.ColMap("ReasonId").SetNotNull(true)

}

type UserCredit struct {
	Id     int `json:"id"     db:"id"       `
	UserId int `json:"user"   db:"user_id"     `
	Credit int `json:"credit" db:"credit"   `
}

func (u UserCredit) setUpTable(db *gorp.DbMap) {
	tbl := db.AddTableWithName(u, "user_credit")
	tbl.SetKeys(true, "Id")
	tbl.ColMap("Id").SetNotNull(true)
	tbl.ColMap("UserId").SetNotNull(true)
	tbl.ColMap("Credit").SetNotNull(true)

}

type Cost struct {
	Id     int    `json:"id"        db:"id"       ` //job id
	Name   string `json:"name"      db:"name"     ` //job name
	Amount int    `json:"amount"    db:"amount"   ` //how much does the job cost (if the job is lower than 0, mean's it's income)
}

func (c Cost) setUpTable(db *gorp.DbMap) {
	tbl := db.AddTableWithName(c, "jobs")
	tbl.SetKeys(true, "Id")
	tbl.ColMap("Id").SetNotNull(true)
	tbl.ColMap("Name").SetNotNull(true)
	tbl.ColMap("Amount").SetNotNull(true)

}
