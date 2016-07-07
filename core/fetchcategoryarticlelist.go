package core

import (
	"math"

	"gopkg.in/mgo.v2/bson"

	"git.furqan.io/faqapp/faqapp/data"
	"git.furqan.io/faqapp/faqapp/db"
)

type FetchCategoryArticleList struct {
	CategoryID string

	ArticleStore  db.ArticleStore
	CategoryStore db.CategoryStore
}

func (a FetchCategoryArticleList) Do() (Result, error) {
	cat, err := a.CategoryStore.Get(bson.ObjectIdHex(a.CategoryID))
	if err != nil {
		return nil, DatabaseError{"CreateArticle", err}
	}

	arts, err := a.ArticleStore.ListCategory(cat.ID, 0, math.MaxInt32)
	if err != nil {
		return nil, DatabaseError{"FetchCategoryArticleList", err}
	}
	return FetchCategoryArticleListRes{
		Articles: arts,
	}, nil
}

type FetchCategoryArticleListRes struct {
	Articles []data.Article
}
