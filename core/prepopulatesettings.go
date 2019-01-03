package core

import (
	"git.furqansoftware.net/faqapp/faqapp/data"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

var DefaultSettings = map[string]interface{}{}

type PrepopulateSettings struct {
	SettingStore db.SettingStore
}

func (p PrepopulateSettings) Do() (Result, error) {
	for k, v := range DefaultSettings {
		stt, err := p.SettingStore.Get(k)
		if err != nil {
			stt = &data.Setting{
				Key:   k,
				Value: v,
			}
			err = p.SettingStore.Put(stt)
			if err != nil {
				return nil, err
			}
		}
	}
	return PrepopulateSettingsRes{}, nil
}

type PrepopulateSettingsRes struct {
	Settings []data.Setting
}
