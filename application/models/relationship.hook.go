package models

import (
	"time"

	"github.com/Thenecromance/OurStories/utility/id"
	"gopkg.in/gorp.v2"
)

func (r *Relationship) PreInsert(s gorp.SqlExecutor) error {
	r.RelationId = id.Generate()
	r.AssociateTime = time.Now().UnixMilli()
	return nil
}

func (r *Relationship) PostInsert(s gorp.SqlExecutor) error {
	s.Insert(&RelationShipHistory{
		UserID:        r.UserID,
		OperationType: Binding,
		Operation:     r.RelationType,
		OperationTime: r.AssociateTime,
		ReceiverID:    r.FriendID,
		OperationUser: r.UserID,
	})

	return nil
}

func (rh *RelationShipHistory) PreInsert(s gorp.SqlExecutor) error {
	rh.ID = id.Generate()

	return nil
}
