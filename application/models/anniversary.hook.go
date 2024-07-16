package models

import (
	"encoding/json"
	"github.com/Thenecromance/OurStories/utility/id"
	"gopkg.in/gorp.v2"
	"time"
)

//------------------------------------------------------------
// hooks
//------------------------------------------------------------

func (a *Anniversary) PreInsert(s gorp.SqlExecutor) error {
	a.Id = id.Generate()
	a.CreatedTime = time.Now().UnixMilli()

	buf, err := json.Marshal(a.SharedWith)
	if err != nil {
		a.SharedWithMarshaled = ""
	}
	a.SharedWithMarshaled = string(buf)
	return nil
}

func (a *Anniversary) PostGet(s gorp.SqlExecutor) error {
	err := json.Unmarshal([]byte(a.SharedWithMarshaled), &a.SharedWith)
	return err
}
