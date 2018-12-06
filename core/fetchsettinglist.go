package core

import (
	"math"

	"git.furqansoftware.net/faqapp/faqapp/data"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

type FetchSettingList struct {
	SettingStore db.SettingStore
}

func (a FetchSettingList) Do() (Result, error) {
	stts, err := a.SettingStore.List(0, math.MaxInt32)
	if err != nil {
		return nil, DatabaseError{"FetchSettingList", err}
	}
	return FetchSettingListRes{
		Settings: stts,
	}, nil
}

type FetchSettingListRes struct {
	Settings []data.Setting
}
