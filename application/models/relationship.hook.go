package models

import (
	"github.com/Thenecromance/OurStories/utility/id"
	"gopkg.in/gorp.v2"
)

func (r *Relationship) PreInsert(s gorp.SqlExecutor) error {
	r.RelationId = id.Generate()

	return nil
}
