package anniversary

import (
	"github.com/Thenecromance/OurStories/backend/anniversary/data"
	"github.com/Thenecromance/OurStories/base/SQL"
	"github.com/Thenecromance/OurStories/base/logger"
	"gopkg.in/gorp.v2"
	"time"
)

type ResponseAnniversary struct {
	data.Anniversary
	TotalSpend int `json:"total_spend"`  // this filed will be calculated by the server until now
	TimeToNext int `json:"time_to_next"` // this filed will be calculated by the server until the next anniversary
}

func (r *ResponseAnniversary) calculate() {
	start := time.Date(r.Year, time.Month(r.Month), r.Day, 0, 0, 0, 0, time.Local)
	now := time.Now()
	r.TotalSpend = int(now.Sub(start).Hours() / 24)

	next := time.Date(now.Year(), time.Month(r.Month), r.Day, 0, 0, 0, 0, time.Local)
	if next.Before(now) {
		next = next.AddDate(1, 0, 0)
	}
	r.TimeToNext = int(next.Sub(now).Hours() / 24)
}

type model struct {
	handler *gorp.DbMap
}

func (m *model) init() {
	m.handler = SQL.Default()
	data.Anniversary{}.SetupTable(m.handler)
}
func (m *model) GetAnniversaryList() (result []ResponseAnniversary) {
	anniversarys, err := m.handler.Select(data.Anniversary{}, "select * from anniversary")
	if err != nil {
		logger.Get().Errorf("failed to get anniversary list with error: %s", err.Error())
		return nil
	}

	result = make([]ResponseAnniversary, 0, len(anniversarys))
	for _, v := range anniversarys {
		ani := ResponseAnniversary{Anniversary: v.(data.Anniversary)}
		ani.calculate()
		result = append(result, ani)
	}
	return
}

func newModel() *model {
	m := &model{}
	return m
}

func Test() {
	a := ResponseAnniversary{
		Anniversary: data.Anniversary{
			Year:  2021,
			Month: 1,
			Day:   1,
		},
	}
	a.calculate()
	logger.Get().Infof("total spend: %d %d", a.TotalSpend, a.TimeToNext)

}
