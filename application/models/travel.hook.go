package models

import (
	"encoding/json"
	"github.com/Thenecromance/OurStories/utility/id"
	"github.com/Thenecromance/OurStories/utility/log"
	"gopkg.in/gorp.v2"
)

func (t *Travel) PreInsert(s gorp.SqlExecutor) error {
	t.Id = id.Generate()
	bytes, err := json.Marshal(t.TogetherWith)
	if err != nil {
		return err
	}
	t.TogetherWithMarshaled = string(bytes)
	return nil
}

func (t *Travel) PostInsert(s gorp.SqlExecutor) error {
	err := s.Insert(&TravelLog{
		TravelId: t.Id,
	})
	if err != nil {
		log.Error("error in inserting travel log", err)
		return err
	}
	return nil
}
