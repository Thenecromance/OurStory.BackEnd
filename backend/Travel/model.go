package Travel

import (
	"errors"
	"github.com/Thenecromance/OurStories/base/SQL"
	"github.com/Thenecromance/OurStories/base/logger"
	"gopkg.in/gorp.v2"
	"strconv"
)

const (
	Prepare = iota + 1
	Traveling
	Finished
)

// Client side data struct
type Data struct {
	Id           int    `json:"id"          `
	State        int    `json:"state"         form:"state"    binding:"required"`
	UserId       int    `json:"owner"         form:"owner"    binding:"required"` // the user who create this travel
	StartTime    int64  `json:"start"         form:"start"    binding:"required"` // this stamp is the time when the travel start, so it's need to be required
	EndTime      int64  `json:"end"           form:"end"    binding:"required"`   // this stamp stored the time when the travel end
	Location     string `json:"location"      form:"location" binding:"required"` // the location where he/she go
	TogetherWith []int  `json:"together"      form:"together"`                    //don't need to required, maybe go alone
	Logs         string `json:"logs"          form:"logs"`                        //if the travel is prepare travel, nothing will be here
	ImagePath    string `json:"img"           form:"img" `                        // I hate stored image in database, so I will store the path of image
}

type UpdateData struct {
	Id           int    `json:"id"              form:"id"   binding:"required"`
	State        int    `json:"state"           form:"state"    `
	UserId       int    `json:"owner"           form:"owner"    ` // the user who create this travel
	Start        int64  `json:"start"           form:"start"    ` // this stamp is the time when the travel start, so it's need to be required
	End          int64  `json:"end"             form:"end"    `   // this stamp stored the time when the travel end
	Location     string `json:"location"        form:"location" ` // the location where he/she go
	TogetherWith []int  `json:"together"        form:"together"`  //don't need to required, maybe go alone
	Logs         string `json:"logs"            form:"logs"`      //if the travel is prepare travel, nothing will be here
	ImagePath    string `json:"img"             form:"img" `      // I hate stored image in database, so I will store the path of image
}

// DbData due to the mysql does not support the array type, so I need to store the array as string
type DbData struct {
	Id           int    `json:"id"          db:"id"`
	State        int    `json:"state"       db:"state"      form:"state"    binding:"required"`
	UserId       int    `json:"owner"       db:"owner"      form:"owner"    binding:"required"` // the user who create this travel
	StartStamp   int64  `json:"start"       db:"start"      form:"start"    binding:"required"` // this stamp is the time when the travel start, so it's need to be required
	EndStamp     int64  `json:"end"         db:"end"        form:"end"      binding:"required"` // this stamp stored the time when the travel end
	Location     string `json:"location"    db:"location"   form:"location" binding:"required"` // the location where he/she go
	TogetherWith string `json:"together"    db:"together"   form:"together"`                    //don't need to required, maybe go alone
	Logs         string `json:"logs"        db:"logs"       form:"logs"`                        //if the travel is prepare travel, nothing will be here
	ImagePath    string `json:"img"         db:"img"        form:"img" `                        // I hate stored image in database, so I will store the path of image
}

// binding the table with gorp
func (d DbData) setUpTable(db *gorp.DbMap) error {
	logger.Get().Info("start to binding Data with table travel")
	tbl := db.AddTableWithName(d, "travel")
	tbl.SetKeys(true, "Id")
	tbl.ColMap("State").SetNotNull(true)
	tbl.ColMap("UserId").SetNotNull(true)
	tbl.ColMap("Location").SetNotNull(true)
	tbl.ColMap("StartStamp").SetNotNull(true)

	return db.CreateTablesIfNotExists()
}

// transfer the data to dbData which could be stored in database (2 different only in the togetherWith)
func transferDataToDbData(data Data) DbData {
	d := DbData{
		Id:           data.Id,
		State:        data.State,
		UserId:       data.UserId,
		StartStamp:   data.StartTime,
		EndStamp:     data.EndTime,
		Location:     data.Location,
		TogetherWith: "",
		Logs:         data.Logs,
		ImagePath:    data.ImagePath,
	}
	if len(data.TogetherWith) > 0 {
		for _, v := range data.TogetherWith {
			d.TogetherWith += strconv.Itoa(v) + ","
		}
	}
	return d
}

// transfer the dbData to data which could be used in the code
func fromDbDataToData(data *DbData) Data {
	d := Data{
		Id:           data.Id,
		State:        data.State,
		UserId:       data.UserId,
		StartTime:    data.StartStamp,
		EndTime:      data.EndStamp,
		Location:     data.Location,
		TogetherWith: []int{},
		Logs:         data.Logs,
		ImagePath:    data.ImagePath,
	}
	if data.TogetherWith != "" {
		ids := data.TogetherWith
		id := ""
		for _, v := range ids {
			if v == ',' {
				i, _ := strconv.Atoi(id)
				d.TogetherWith = append(d.TogetherWith, i)
				id = ""
			} else {
				id += string(v)
			}
		}
	}

	return d
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
	DbData{}.setUpTable(m.db)
	return nil
}
func (m *Model) AddTravel(data Data) error {
	err := m.initdb()
	if err != nil {
		return err
	}
	dbData := transferDataToDbData(data)
	err = m.db.Insert(&dbData) //insert data into database
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) getDbDataById(id int) (*DbData, error) {
	err := m.initdb()
	if err != nil {
		return nil, err
	}
	data, err := m.db.Get(DbData{}, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New("no data found")
	}
	return data.(*DbData), nil
}

func (m *Model) GetTravelByUserId(id int) ([]Data, error) {
	err := m.initdb()
	if err != nil {
		return nil, err
	}
	datas, err := m.db.Select(DbData{}, "SELECT * FROM `travel` WHERE `owner` = ? ORDER BY `start` desc ", id)
	if err != nil {
		return nil, err
	}

	var datas2 []Data
	for _, v := range datas {
		datas2 = append(datas2, fromDbDataToData(v.(*DbData)))
	}

	return datas2, nil
}
func (m *Model) GetTravelById(id int) (Data, error) {
	err := m.initdb()
	if err != nil {
		return Data{}, err
	}
	data, err := m.db.Get(DbData{}, id)
	if err != nil {
		return Data{}, err
	}
	return fromDbDataToData(data.(*DbData)), nil
}

func (m *Model) updateToDatabase(data UpdateData) error {
	err := m.initdb()
	if err != nil {
		return err
	}
	logger.Get().Info(data.Id)
	//request the old data
	old, err := m.getDbDataById(data.Id)
	if err != nil {
		return err
	}
	//transfer the data to dbData
	dbData := transferDataToDbData(Data{
		Id:           data.Id,
		State:        data.State,
		UserId:       data.UserId,
		StartTime:    data.Start,
		EndTime:      data.End,
		Location:     data.Location,
		TogetherWith: data.TogetherWith,
		Logs:         data.Logs,
		ImagePath:    data.ImagePath,
	})
	logger.Get().Infof("old data:%v", old)

	{
		if dbData.State == 0 {
			dbData.State = old.State
		}
		if dbData.UserId == 0 {
			dbData.UserId = old.UserId
		}
		if dbData.StartStamp == 0 {
			dbData.StartStamp = old.StartStamp
		}
		if dbData.Location == "" {
			dbData.Location = old.Location
		}
		if dbData.TogetherWith == "" {
			dbData.TogetherWith = old.TogetherWith
		}
		if dbData.Logs == "" {
			dbData.Logs = old.Logs
		}
		if dbData.ImagePath == "" {
			dbData.ImagePath = old.ImagePath
		}
	}
	logger.Get().Infof("old data:%v", dbData)

	_, err = m.db.Update(&dbData)
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) RemoveTravel(id int) error {
	err := m.initdb()
	if err != nil {
		return err
	}

	_, err = m.db.Delete(&DbData{
		Id: id,
	})
	if err != nil {
		return err
	}
	return nil
}
