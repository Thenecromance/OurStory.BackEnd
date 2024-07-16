package models

import (
	"encoding/json"
	"time"

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
		TravelId:   t.Id,
		ModifiedBy: t.UserId,
		Message:    "Travel created",
	})
	if err != nil {
		log.Error("error in inserting travel log", err)
		return err
	}
	return nil
}

func (t *Travel) PostUpdate(s gorp.SqlExecutor) error {
	err := s.Insert(&TravelLog{
		TravelId:   t.Id,
		ModifiedBy: t.UserId,
		Message:    "Travel updated",
	})
	if err != nil {
		log.Error("error in inserting travel log", err)
		return err
	}
	return nil
}

func (t *Travel) PostDelete(s gorp.SqlExecutor) error {
	err := s.Insert(&TravelLog{
		TravelId:   t.Id,
		ModifiedBy: t.UserId,
		Message:    "Travel deleted",
	})
	if err != nil {
		log.Error("error in inserting travel log", err)
		return err
	}
	return nil
}

func (t *TravelLog) PreInsert(s gorp.SqlExecutor) error {
	t.LogId = id.Generate()
	t.ModifiedAt = time.Now().UnixMilli()
	return nil
}
