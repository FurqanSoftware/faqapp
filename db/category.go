package db

import (
	"git.furqansoftware.net/faqapp/faqapp/data"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const CategoryC = "categories"

type CategoryStore struct {
	Session *Session
}

func (s CategoryStore) Get(id bson.ObjectId) (*data.Category, error) {
	cat := data.Category{}
	err := s.Session.DB("").C(CategoryC).
		FindId(id).
		One(&cat)
	if err == mgo.ErrNotFound {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &cat, nil
}

func (s CategoryStore) GetBySlug(slug string) (*data.Category, error) {
	cat := data.Category{}
	err := s.Session.DB("").C(CategoryC).
		Find(bson.M{"slug": slug}).
		One(&cat)
	if err == mgo.ErrNotFound {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &cat, nil
}

func (s CategoryStore) List(skip, limit int) ([]data.Category, error) {
	cats := []data.Category{}
	err := s.Session.DB("").C(CategoryC).
		Find(nil).
		Sort("order").
		Skip(skip).
		Limit(limit).
		All(&cats)
	if err != nil {
		return nil, err
	}
	return cats, nil
}

func (s CategoryStore) Put(cat *data.Category) error {
	return put(s.Session, CategoryC, cat, cat.ID)
}
