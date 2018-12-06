package db

import (
	"git.furqansoftware.net/faqapp/faqapp/data"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const SettingC = "settings"

type SettingStore struct {
	Session *Session
}

func (s SettingStore) Get(key string) (*data.Setting, error) {
	stt := data.Setting{}
	err := s.Session.DB("").C(SettingC).
		Find(bson.M{"key": key}).
		One(&stt)
	if err == mgo.ErrNotFound {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &stt, nil
}

func (s SettingStore) List(skip, limit int) ([]data.Setting, error) {
	stts := []data.Setting{}
	err := s.Session.DB("").C(SettingC).
		Find(nil).
		Skip(skip).
		Limit(limit).
		All(&stts)
	if err != nil {
		return nil, err
	}
	return stts, nil
}

func (s SettingStore) Put(stt *data.Setting) error {
	_, err := s.Session.DB("").C(SettingC).Upsert(bson.M{"key": stt.Key}, stt)
	return err
}
