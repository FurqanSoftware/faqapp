package core

import (
	"math"

	"gopkg.in/mgo.v2/bson"

	"git.furqan.io/faqapp/faqapp/data"
	"git.furqan.io/faqapp/faqapp/db"
)

type FetchArticleList struct {
	CategoryID string

	ArticleStore  db.ArticleStore
	CategoryStore db.CategoryStore
}

func (a FetchArticleList) Do() (Result, error) {
	cat, err := a.CategoryStore.Get(bson.ObjectIdHex(a.CategoryID))
	if err != nil {
		return nil, DatabaseError{"CreateArticle", err}
	}

	arts, err := a.ArticleStore.List(cat.ID, 0, math.MaxInt32)
	if err != nil {
		return nil, DatabaseError{"FetchArticleList", err}
	}
	return FetchArticleListRes{
		Articles: arts,
	}, nil
}

type FetchArticleListRes struct {
	Articles []data.Article
}
