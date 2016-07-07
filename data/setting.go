package data

import "encoding/json"

type Setting struct {
	Key   string      `bson:"key"`
	Value interface{} `bson:"value"`
}

func (s Setting) ValueJSON() (json.RawMessage, error) {
	v, err := json.Marshal(s.Value)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(v), nil
}

func (s Setting) ValueString() string {
	v, _ := s.Value.(string)
	return v
}
