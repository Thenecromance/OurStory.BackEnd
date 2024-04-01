package User

import (
	"github.com/Thenecromance/OurStories/base/SQL"
	uuid "github.com/satori/go.uuid"
)

func BindRelationsTable(r Relations) {
	tbl := SQL.Default().AddTableWithName(r, "relations")
	tbl.SetKeys(true, "Id")
	SQL.Default().CreateTablesIfNotExists()
}

// CreateDefaultRelation this method should only called when the first time user register , it will called once time for each user
func CreateDefaultRelation(userId int) Relations {
	return Relations{
		Mate1:    userId,
		Mate2:    0,
		LinkCode: uuid.NewV4().String(), // using UUID to generate the link code
	}
}

type Relations struct {
	Id          int    `db:"id"             json:"id"`          // self id
	Mate1       int    `db:"mate1"          json:"mate1"`       // mate1 id
	Mate2       int    `db:"mate2"          json:"mate2"`       // mate2 id
	LinkCode    string `db:"link_code"      json:"link_code"`   // the code to link two user
	LinkDate    int64  `db:"link_date"      json:"link_date"`   // the date when two user link
	Anniversary int64  `db:"anniversary"    json:"anniversary"` // the date when two user link
}

func (r *Relations) InsertToSQL() error {
	return SQL.Default().Insert(r)
}

func (r *Relations) removeLinkCode() error {
	r.LinkCode = ""
	_, err := SQL.Default().Update(r)
	return err
}

func (r *Relations) UpdateRelation() error {
	return nil
}
func (r *Relations) GetRelationsByUserId(id int) []Relations {
	var res []Relations
	_, err := SQL.Default().Select(&res, "select * from relations where mate1 = ? or mate2 = ?", id, id)
	if err != nil {
		return nil
	}
	return res
}
