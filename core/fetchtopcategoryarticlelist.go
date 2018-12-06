package core

import (
	"gopkg.in/mgo.v2/bson"

	"git.furqansoftware.net/faqapp/faqapp/data"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

type FetchTopCategoryArticleList struct {
	CategoryID string

	ArticleStore  db.ArticleStore
	CategoryStore db.CategoryStore
}

func (a FetchTopCategoryArticleList) Validate() error {
	if !bson.IsObjectIdHex(a.CategoryID) {
		return ValidationError{"FetchTopCategoryArticleList", "CategoryID", IssueInvalid}
	}
	return nil
}

func (a FetchTopCategoryArticleList) Do() (Result, error) {
	cat, err := a.CategoryStore.Get(bson.ObjectIdHex(a.CategoryID))
	if err != nil {
		return nil, DatabaseError{"FetchTopCategoryArticleList", err}
	}

	arts, err := a.ArticleStore.ListCategory(cat.ID, 0, 8)
	if err != nil {
		return nil, DatabaseError{"FetchTopCategoryArticleList", err}
	}
	return FetchTopCategoryArticleListRes{
		Articles: arts,
	}, nil
}

type FetchTopCategoryArticleListRes struct {
	Articles []data.Article
}
