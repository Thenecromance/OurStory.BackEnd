package anniversary

import (
	"github.com/Thenecromance/OurStories/utility/SQL"
	"github.com/Thenecromance/OurStories/utility/log"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/gorp.v2"
	"time"
)

type ResponseAnniversary struct {
	data.Anniversary
	TotalSpend int `json:"total_spend"`  // this filed will be calculated by the server until now
	TimeToNext int `json:"time_to_next"` // this filed will be calculated by the server until the next anniversary
}

func (r *ResponseAnniversary) calculate() {
	start := time.Unix(r.TimeStamp, 0)
	now := time.Now()
	r.TotalSpend = int(now.Sub(start).Hours() / 24)

	_, month, day := start.Date()
	next := time.Date(now.Year(), month, day, 0, 0, 0, 0, time.Local)
	if next.Before(now) {
		next = next.AddDate(1, 0, 0)
	}
	r.TimeToNext = int(next.Sub(now).Hours() / 24)

}

type Model struct {
	handler *gorp.DbMap
}

func (m *Model) init() {
	log.Info("start to init anniversary Model")

	m.handler = SQL.Default()
	data.Anniversary{}.SetupTable(m.handler)

	log.Info("init anniversary Model success")
}
func (m *Model) GetAnniversaryList() (result []ResponseAnniversary) {
	objects, err := m.handler.Select(data.Anniversary{}, "select * from anniversary")
	if err != nil {
		log.Errorf("failed to get anniversary list with error: %s", err.Error())
		return nil
	}

	result = make([]ResponseAnniversary, 0, len(objects))
	for _, v := range objects {
		vv, _ := v.(*data.Anniversary)
		ani := ResponseAnniversary{Anniversary: *vv}
		ani.calculate()
		result = append(result, ani)
	}
	return
}

func (m *Model) AddAnniversary(ani data.Anniversary) error {
	ani.Id = uuid.NewV4().String()
	err := m.handler.Insert(&ani)
	if err != nil {
		log.Errorf("failed to insert anniversary with error: %s", err.Error())
		return err
	}
	return nil
}

// GetAnniversaryById get anniversary by uid from database
func (m *Model) GetAnniversaryById(id string) *data.Anniversary {
	anni, err := m.handler.Get(data.Anniversary{}, id)
	if err != nil {
		return nil
	}
	return anni.(*data.Anniversary)
}

func NewModel() *Model {
	m := &Model{}
	m.init()
	return m
}
