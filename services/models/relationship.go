package models

type Relationship struct {
	UserID        int   `json:"user_id,omitempty" db:"user_id" `
	FriendID      int   `json:"friend_id,omitempty" db:"friend_id"`
	RelationType  int   `json:"relation_type,omitempty" db:"relation_type"`
	AssociateTime int64 `json:"associate_time,omitempty" db:"associate_time"`
}

const (
	Unknown = iota // this is the default value
	Friend         // means the user is a friend . also the users should be associated with more than 1 user
	Couple         // means the user is a couple . also the users should be associated with only 1 user

)

type RelationShipResponse struct {
	URL          string `json:"url,omitempty"`
	RelationType int    `json:"relation_type,omitempty"` // identify the relation type
}
