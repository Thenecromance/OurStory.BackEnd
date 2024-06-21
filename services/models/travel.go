package models

type TravelInfo struct {
	Id           string `json:"id"          db:"id"`
	State        int    `json:"state"       db:"state"        form:"state"`
	UserId       int    `json:"owner"       db:"owner"        form:"owner"     binding:"required"` // the user who create this travel
	StartTime    int64  `json:"start"       db:"start"        form:"start"     binding:"required"` // this stamp is the time when the travel start, so it's need to be required
	EndTime      int64  `json:"end"         db:"end"          form:"end"       binding:"required"` // this stamp stored the time when the travel end
	Location     string `json:"location"    db:"location"     form:"location"  binding:"required"` // the location where he/she go
	Details      string `json:"details"        db:"details"   form:"details"`                      //if the travel is prepare travel, nothing will be here
	TogetherWith []int  `json:"together"      form:"together"`                                     //don't need to required, maybe go alone
	//ImagePath string `json:"img"           form:"img" ` // I hate stored image in database, so I will store the path of image
}
