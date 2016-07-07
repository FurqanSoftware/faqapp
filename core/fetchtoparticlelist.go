package core

import (
	"gopkg.in/mgo.v2/bson"

	"git.furqan.io/faqapp/faqapp/data"
	"git.furqan.io/faqapp/faqapp/db"
)

type FetchTopArticleList struct {
	CategoryID string

	ArticleStore db.ArticleStore
}

func (a FetchTopArticleList) Validate() error {
	if !bson.IsObjectIdHex(a.CategoryID) {
		return ValidationError{"FetchTopArticleList", "CategoryID", IssueInvalid}
	}
	return nil
}

func (a FetchTopArticleList) Do() (Result, error) {
	arts, err := a.ArticleStore.ListCategory(bson.ObjectIdHex(a.CategoryID), 0, 8)
	if err != nil {
		return nil, DatabaseError{"FetchTopArticleList", err}
	}
	return FetchTopArticleListRes{
		Articles: arts,
	}, nil
}

type FetchTopArticleListRes struct {
	Articles []data.Article
}
