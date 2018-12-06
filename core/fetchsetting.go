package core

import (
	"git.furqansoftware.net/faqapp/faqapp/data"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

type FetchSetting struct {
	Key string

	SettingStore db.SettingStore
}

func (a FetchSetting) Do() (Result, error) {
	stt, err := a.SettingStore.Get(a.Key)
	if err != nil && err != db.ErrNotFound {
		return nil, DatabaseError{"FetchSetting", err}
	}
	return FetchSettingRes{
		Setting: stt,
	}, nil
}

type FetchSettingRes struct {
	Setting *data.Setting
}
