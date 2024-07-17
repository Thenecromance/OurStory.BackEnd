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
	Id                    int64   `json:"travel_id"       redis:"id"                       db:"travel_id"                         `
	State                 int     `json:"state"           redis:"state"                    db:"state"          form:"state"       `
	UserId                int64   `json:"user_id"         redis:"user_id"                  db:"user_id"        form:"user_id"          binding:"required"   `  // the user who create this travel
	StartTime             int64   `json:"travel_start"    redis:"start_time"               db:"travel_start"   form:"travel_start"     binding:"required"   `  // this stamp is the time when the travel start, so it's need to be required
	EndTime               int64   `json:"travel_end"      redis:"end_time"                 db:"travel_end"     form:"travel_end"       binding:"required"    ` // this stamp stored the time when the travel end
	Location              string  `json:"location"        redis:"location"                 db:"location"       form:"location"         binding:"required"   `  // the location where he/she go
	Detail                string  `json:"detail"          redis:"detail"                   db:"detail"         form:"detail"    `                              //if the travel is prepare travel, nothing will be here
	TogetherWithMarshaled string  `json:"-"               redis:"-"                        db:"together"       form:"-"         `                              // this is the user list who will go with the owner
	ImagePath             string  `json:"image"           redis:"image_path"               db:"image"          form:"img"           `                          // I hate stored image in database, so I will store the path of image
	CreateAt              int64   `json:"created_at"      redis:"create_at"                db:"created_at"                          `                          // the time when the travel is created
	TogetherWith          []int64 `json:"together"        redis:"together_with"            db:"-"              form:"together"          `                      // this is the user list who will go with the owner
}

type TravelLog struct {
	LogId      int64  `db:"log_id"`
	TravelId   int64  `db:"travel_id"`
	ModifiedBy int64  `db:"modified_by"`
	ModifiedAt int64  `db:"modified_at"`
	Message    string `db:"message"`
}

// =======================================================
// Data Transfer Object
// =======================================================

// UserTravelInfo is the struct that will be used to store the user travel information
type UserTravelInfo struct {
}
