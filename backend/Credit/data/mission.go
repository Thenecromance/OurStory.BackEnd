package data

type Mission struct {
	Id          int    `json:"id"`          // Mission ID
	Name        string `json:"name"`        // Mission name
	Description string `json:"description"` // Mission description
	Reward      int    `json:"reward"`      // When the job is finished, the reward will add to the user's balance
	LimitCount  int    `json:"limit_count"` // each account can only do this mission for a limited number of times if this value is -1 means this mission has no limit
	CoolDown    int    `json:"cool_down"`   // The time limit for the mission if this value is -1 means this mission has no time limit
	Enable      bool   `json:"enable"`      // Whether the mission is enabled
}

// Store the user's mission history
type MissionHistory struct {
	UserId    int `json:"user_id"`    // who did this mission
	MissionId int `json:"mission_id"` // which mission did the user do
	Times     int `json:"times"`      // how many times did the user do this mission
	Reward    int `json:"reward"`     // how many rewards did the user get
}
