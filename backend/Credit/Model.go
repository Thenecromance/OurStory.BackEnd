package Credit

import (
	"github.com/Thenecromance/OurStories/base/SQL"
	"gopkg.in/gorp.v2"
)

type Model struct {
	db   *gorp.DbMap
	jobs []Cost
}

func (m *Model) initdb() error {
	if m.db != nil {
		return nil
	}
	m.db = SQL.Default()
	//create Info Table
	Modified{}.setUpTable(m.db)
	UserCredit{}.setUpTable(m.db)
	Cost{}.setUpTable(m.db)
	return nil
}
