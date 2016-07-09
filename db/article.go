package db

import (
	"git.furqan.io/faqapp/faqapp/data"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const ArticleC = "articles"

type ArticleStore struct {
	Session *Session
}

func (s ArticleStore) Get(id bson.ObjectId) (*data.Article, error) {
	art := data.Article{}
	err := s.Session.DB("").C(ArticleC).
		FindId(id).
		One(&art)
	if err == mgo.ErrNotFound {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &art, nil
}

func (s ArticleStore) GetBySlug(catID bson.ObjectId, slug string) (*data.Article, error) {
	art := data.Article{}
	err := s.Session.DB("").C(ArticleC).
		Find(bson.M{"category_id": catID, "slug": slug}).
		One(&art)
	if err == mgo.ErrNotFound {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &art, nil
}

func (s ArticleStore) List(skip, limit int) ([]data.Article, error) {
	arts := []data.Article{}
	err := s.Session.DB("").C(ArticleC).
		Find(nil).
		Skip(skip).
		Limit(limit).
		All(&arts)
	if err != nil {
		return nil, err
	}
	return arts, nil
}

func (s ArticleStore) ListCategory(catID bson.ObjectId, skip, limit int) ([]data.Article, error) {
	arts := []data.Article{}
	err := s.Session.DB("").C(ArticleC).
		Find(bson.M{"category_id": catID}).
		Sort("order").
		Skip(skip).
		Limit(limit).
		All(&arts)
	if err != nil {
		return nil, err
	}
	return arts, nil
}

func (s ArticleStore) Put(art *data.Article) error {
	callPutHooks(art, art.ID == "")
	return put(s.Session, ArticleC, art, art.ID)
}
