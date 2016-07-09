package core

import (
	"encoding/json"

	"git.furqan.io/faqapp/faqapp/data"
	"git.furqan.io/faqapp/faqapp/db"
)

type UpdateSetting struct {
	Key   string
	Value string

	SettingStore db.SettingStore
}

func (a UpdateSetting) Validate() error {
	if a.Key == "" {
		return ValidationError{"UpdateSetting", "Key", IssueMissing}
	}
	if len(a.Key) > 128 {
		return ValidationError{"UpdateSetting", "Key", IssueTooLong}
	}
	return nil
}

func (a UpdateSetting) Do() (Result, error) {
	stt, err := a.SettingStore.Get(a.Key)
	if err != nil && err != db.ErrNotFound {
		return nil, DatabaseError{"UpdateSetting", err}
	}
	if stt == nil {
		stt = &data.Setting{
			Key: a.Key,
		}
	}

	var v interface{}
	err = json.Unmarshal([]byte(a.Value), &v)
	if err != nil {
		return nil, err
	}

	stt.Value = v
	err = a.SettingStore.Put(stt)
	if err != nil {
		return nil, DatabaseError{"UpdateSetting", err}
	}

	return UpdateSettingRes{
		Setting: stt,
	}, nil
}

type UpdateSettingRes struct {
	Setting *data.Setting
}
