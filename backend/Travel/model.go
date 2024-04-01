package Travel

import (
	"github.com/Thenecromance/OurStories/base/SQL"
	"github.com/Thenecromance/OurStories/base/logger"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/gorp.v2"
)

const (
	Prepare = iota + 1
	Ongoing
	Finished
)

// ClientData Client side data struct
type ClientData struct {
	Id        string `json:"id"          db:"id"`
	State     int    `json:"state"       db:"state"        form:"state"`
	UserId    int    `json:"owner"       db:"owner"        form:"owner"     binding:"required"` // the user who create this travel
	StartTime int64  `json:"start"       db:"start"        form:"start"     binding:"required"` // this stamp is the time when the travel start, so it's need to be required
	EndTime   int64  `json:"end"         db:"end"          form:"end"       binding:"required"` // this stamp stored the time when the travel end
	Location  string `json:"location"    db:"location"     form:"location"  binding:"required"` // the location where he/she go
	Details   string `json:"details"        db:"details"   form:"details"`                      //if the travel is prepare travel, nothing will be here
	//TogetherWith []int  `json:"together"      form:"together"`                    //don't need to required, maybe go alone
	//ImagePath string `json:"img"           form:"img" ` // I hate stored image in database, so I will store the path of image
}

// binding the table with gorp
func (d ClientData) setUpTable(db *gorp.DbMap) error {
	logger.Get().Info("start to binding ClientData with table travel")
	tbl := db.AddTableWithName(d, "travel")
	tbl.SetKeys(false, "Id") // using snowflake to generate the id
	tbl.ColMap("Id").SetNotNull(true)
	tbl.ColMap("State").SetNotNull(true)
	tbl.ColMap("UserId").SetNotNull(true)
	tbl.ColMap("Location").SetNotNull(true)
	tbl.ColMap("StartTime").SetNotNull(true)
	tbl.ColMap("EndTime").SetNotNull(true)

	return db.CreateTablesIfNotExists()
}

type Model struct {
	db *gorp.DbMap
}

func (m *Model) initdb() error {
	if m.db != nil {
		return nil
	}
	m.db = SQL.Default()
	//create Info Table
	ClientData{}.setUpTable(m.db)
	return nil
}

func (m *Model) AddToSQL(data *ClientData) error {
	err := m.initdb()
	if err != nil {
		logger.Get().Error(err)
		return err
	}

	data.Id = uuid.NewV4().String()
	err = m.db.Insert(data) //insert data into database
	if err != nil {
		logger.Get().Error(err)
		return err
	}

	return nil
}

func (m *Model) RemoveTravel(id string) error {
	err := m.initdb()
	logger.Get().Infof("start to remove travel %d", id)
	if err != nil {
		return err
	}
	_, err = m.db.Delete(&ClientData{Id: id})
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) UpdateTravel(data *ClientData) error {
	err := m.initdb()
	if err != nil {
		return err
	}
	_, err = m.db.Update(&data)
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) GetTravelByTravelId(travelId string) (*ClientData, error) {
	err := m.initdb()
	if err != nil {
		return nil, err
	}
	var data ClientData
	err = m.db.SelectOne(&data, "select * from travel where id=?", travelId)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (m *Model) GetTravelListByUser(userId int) ([]ClientData, error) {
	err := m.initdb()
	if err != nil {
		return nil, err
	}
	var data []ClientData
	_, err = m.db.Select(&data, "select * from travel where owner=? order by start DESC", userId)
	if err != nil {
		return nil, err
	}
	return data, nil
}
