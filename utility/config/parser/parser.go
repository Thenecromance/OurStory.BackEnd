package parser

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
)

type Parser interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

type Json struct {
}

func (j Json) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
func (j Json) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

type Yaml struct {
}

func (y Yaml) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (y Yaml) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
