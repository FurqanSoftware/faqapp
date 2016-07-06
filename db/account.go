package db

import (
	"git.furqan.io/faqapp/faqapp/data"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const AccountC = "accounts"

type AccountStore struct {
	Session *Session
}

func (s AccountStore) Get(id bson.ObjectId) (*data.Account, error) {
	acc := data.Account{}
	err := s.Session.DB("").C(AccountC).
		FindId(id).
		One(&acc)
	if err == mgo.ErrNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &acc, nil
}

func (s AccountStore) GetByHandle(handle string) (*data.Account, error) {
	acc := data.Account{}
	err := s.Session.DB("").C(AccountC).
		Find(bson.M{"handle": handle}).
		One(&acc)
	if err == mgo.ErrNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &acc, nil
}

func (s AccountStore) List(skip, limit int) ([]data.Account, error) {
	accs := []data.Account{}
	err := s.Session.DB("").C(AccountC).
		Find(nil).
		Skip(skip).
		Limit(limit).
		All(&accs)
	if err != nil {
		return nil, err
	}
	return accs, nil
}

func (s AccountStore) Put(acc *data.Account) error {
	callPutHooks(acc, acc.ID == "")
	return put(s.Session, AccountC, acc, acc.ID)
}
