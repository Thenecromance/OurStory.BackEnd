package models

import (
	"github.com/Thenecromance/OurStories/utility/id"
	"gopkg.in/gorp.v2"
)

// Relationship is the struct that defines the relationship between users which stores in the database
type Relationship struct {
	RelationId    int64  `json:"relation_id" db:"relation_id"`
	UserID        int64  `json:"user_id" db:"user_id" `            // the user id
	FriendID      int64  `json:"friend_id" db:"friend_id"`         // associate with the user id
	RelationType  int    `json:"relation_type" db:"relation_type"` // two of the user's relationship type
	Status        string `json:"status" db:"status"`               // the status of the relationship
	AssociateTime int64  `json:"stamp" db:"associate_time"`        // the time when the relationship is created
}

const (
	Unknown = iota // this is the default value
	Friend         // means the user is a friend . also the users should be associated with more than 1 user
	Couple         // means the user is a couple . also the users should be associated with only 1 user

)

// Operations
const (
	Binding      = iota // means the user is binding with other users
	Disassociate        // means the user is disassociate with other users
	Modify              // means the user is modifying the relationship with other users
)

// RelationShipHistory is the history of the user's relationship
// which it will be used to track the user's relationship's operation
// like associate, disassociate
type RelationShipHistory struct {
	ID int64 `json:"id" db:"id"`
	// which user is doing the operation
	UserID int64 `json:"user_id" db:"user_id"`
	// the operation type
	OperationType int `json:"operation_type" db:"operation_type"`
	// when the operation is done
	OperationTime int64 `json:"operation_time" db:"operation_time"`
	// the operation
	Operation int `json:"operation" db:"operation"`
	// the target user id
	ReceiverID int `json:"target_id" db:"target_id"`

	OperationUser int `json:"operation_user" db:"operation_user"`
}

func (r *Relationship) PreInsert(s gorp.SqlExecutor) error {
	r.RelationId = id.Generate()

	return nil
}
