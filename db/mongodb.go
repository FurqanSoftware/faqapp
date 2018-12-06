package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Session struct {
	*mgo.Session
}

type PreCreate interface {
	PreCreate()
}

type PreModify interface {
	PreModify()
}

func Open(url string) (*Session, error) {
	sess, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Session{sess}, nil
}

func callPutHooks(v interface{}, create bool) {
	if create {
		d, ok := v.(PreCreate)
		if ok {
			d.PreCreate()
		}
	} else {
		d, ok := v.(PreModify)
		if ok {
			d.PreModify()
		}
	}
}

func put(sess *Session, col string, v interface{}, id bson.ObjectId) error {
	callPutHooks(v, id == "")
	_, err := sess.DB("").C(col).UpsertId(id, v)
	return err
}
