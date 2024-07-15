package models

const (
	TravelStatePending = iota
	TravelStateOngoing
	TravelStateFinished
)

// =======================================================
// Data In Db
// =======================================================

type Travel struct {
	Id           int    `json:"id"          db:"id"`
	State        int    `json:"state"       db:"state"        form:"state"`
	UserId       int    `json:"owner"       db:"owner"        form:"owner"     binding:"required"` // the user who create this travel
	StartTime    int64  `json:"start"       db:"start"        form:"start"     binding:"required"` // this stamp is the time when the travel start, so it's need to be required
	EndTime      int64  `json:"end"         db:"end"          form:"end"       binding:"required"` // this stamp stored the time when the travel end
	Location     string `json:"location"    db:"location"     form:"location"  binding:"required"` // the location where he/she go
	Details      string `json:"details"     db:"details"      form:"details"`                      //if the travel is prepare travel, nothing will be here
	TogetherWith string `json:"together"    db:"together"     form:"together"`                     // this is the user list who will go with the owner
	ImagePath    string `json:"img"         db:"image"        form:"img" `                         // I hate stored image in database, so I will store the path of image
}

type TravelLog struct {
	LogId      int64  `db:"log_id"`
	TravelId   int    `db:"travel_id"`
	ModifiedBy int    `db:"modified_by"`
	ModifiedAt int64  `db:"modified_at"`
	Message    string `db:"message"`
}

// =======================================================
// Data Transfer Object
// =======================================================

type TravelDTO struct {
	Id           int    `json:"id"          `
	State        int    `json:"state"           form:"state"`
	UserId       int    `json:"owner"           form:"owner"     binding:"required"` // the user who create this travel
	StartTime    int64  `json:"start"           form:"start"     binding:"required"` // this stamp is the time when the travel start, so it's need to be required
	EndTime      int64  `json:"end"             form:"end"       binding:"required"` // this stamp stored the time when the travel end
	Location     string `json:"location"        form:"location"  binding:"required"` // the location where he/she go
	Details      string `json:"details"         form:"details"`                      //if the travel is prepare travel, nothing will be here
	TogetherWith []int  `json:"together"        form:"together"`                     // this is the user list who will go with the owner
	ImagePath    string `json:"img"             form:"img" `                         // I hate stored image in database, so I will store the path of image
}

// UserTravelInfo is the struct that will be used to store the user travel information
type UserTravelInfo struct {
}
