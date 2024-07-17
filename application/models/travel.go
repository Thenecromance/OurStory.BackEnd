package models

import "github.com/Thenecromance/OurStories/application/models/internal"

const (
	TravelStatePending = iota
	TravelStateOngoing
	TravelStateFinished
)

// =======================================================
// Data In Db
// =======================================================

type travelBase struct {
	Id        int64  `json:"travel_id"       redis:"id"                       db:"travel_id"                         `
	State     int    `json:"state"           redis:"state"                    db:"state"          form:"state"       `
	UserId    int64  `json:"user_id"         redis:"user_id"                  db:"user_id"        form:"user_id"          binding:"required"   `  // the user who create this travel
	StartTime int64  `json:"travel_start"    redis:"start_time"               db:"travel_start"   form:"travel_start"     binding:"required"   `  // this stamp is the time when the travel start, so it's need to be required
	EndTime   int64  `json:"travel_end"      redis:"end_time"                 db:"travel_end"     form:"travel_end"       binding:"required"    ` // this stamp stored the time when the travel end
	Location  string `json:"location"        redis:"location"                 db:"location"       form:"location"         binding:"required"   `  // the location where he/she go
	Detail    string `json:"detail"          redis:"detail"                   db:"detail"         form:"detail"    `                              //if the travel is prepare travel, nothing will be here
	ImagePath string `json:"image"           redis:"image_path"               db:"image"          form:"img"           `                          // I hate stored image in database, so I will store the path of image
	CreateAt  int64  `json:"created_at"      redis:"create_at"                db:"created_at"                          `                          // the time when the travel is created
}

type Travel struct {
	travelBase
	TogetherWithMarshaled string  `json:"-"               redis:"-"                        db:"together"       form:"-"         `         // this is the user list who will go with the owner
	TogetherWith          []int64 `json:"together"        redis:"together_with"            db:"-"              form:"together"          ` // this is the user list who will go with the owner
}

func (t *Travel) To() any {
	obj := &internal.TravelToRedis{
		Id:        t.Id,
		State:     t.State,
		UserId:    t.UserId,
		StartTime: t.StartTime,
		EndTime:   t.EndTime,
		Location:  t.Location,
		Detail:    t.Detail,
		ImagePath: t.ImagePath,
		CreateAt:  t.CreateAt,
	}
	for _, v := range t.TogetherWith {
		obj.TogetherWith = append(obj.TogetherWith, byte(v))
	}
	return obj
}
func (t *Travel) From(obj any) {
	redisObj := obj.(*internal.TravelToRedis)
	t.Id = redisObj.Id
	t.State = redisObj.State
	t.UserId = redisObj.UserId
	t.StartTime = redisObj.StartTime
	t.EndTime = redisObj.EndTime
	t.Location = redisObj.Location
	t.Detail = redisObj.Detail
	t.ImagePath = redisObj.ImagePath
	t.CreateAt = redisObj.CreateAt
	for _, v := range redisObj.TogetherWith {
		t.TogetherWith = append(t.TogetherWith, int64(v))
	}
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
