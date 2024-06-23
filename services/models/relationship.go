package models

// Relationship is the struct that defines the relationship between users which stores in the database
type Relationship struct {
	ID            int   `json:"id" db:"id"`
	UserID        int   `json:"user_id" db:"user_id" `              // the user id
	FriendID      int   `json:"friend_id" db:"friend_id"`           // associate with the user id
	RelationType  int   `json:"relation_type" db:"relation_type"`   // two of the user's relationship type
	AssociateTime int64 `json:"associate_time" db:"associate_time"` // the time when the relationship is created
}

const (
	Unknown = iota // this is the default value
	Friend         // means the user is a friend . also the users should be associated with more than 1 user
	Couple         // means the user is a couple . also the users should be associated with only 1 user

)

type RelationShipResponse struct {
	URL          string `json:"url"`
	RelationType int    `json:"relation_type"` // identify the relation type
}

// RelationShipHistory is the history of the user's relationship
// which it will be used to track the user's relationship's operation
// like associate, disassociate
type RelationShipHistory struct {
	// which user is doing the operation
	UserID int `json:"user_id" db:"user_id"`
	// the operation type
	OperationType int `json:"operation_type" db:"operation_type"`
	// when the operation is done
	OperationTime int64 `json:"operation_time" db:"operation_time"`
	// the target user id
	TargetID int `json:"target_id" db:"target_id"`
}
