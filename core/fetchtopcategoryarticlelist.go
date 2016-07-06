package core

import (
	"gopkg.in/mgo.v2/bson"

	"git.furqan.io/faqapp/faqapp/data"
	"git.furqan.io/faqapp/faqapp/db"
)

type FetchTopCategoryArticleList struct {
	CategoryID string

	ArticleStore db.ArticleStore
}

func (a FetchTopCategoryArticleList) Validate() error {
	if !bson.IsObjectIdHex(a.CategoryID) {
		return ValidationError{"FetchTopCategoryArticleList", "CategoryID", IssueInvalid}
	}
	return nil
}

func (a FetchTopCategoryArticleList) Do() (Result, error) {
	arts, err := a.ArticleStore.ListCategoryTop(bson.ObjectIdHex(a.CategoryID), 0, 8)
	if err != nil {
		return nil, DatabaseError{"FetchTopCategoryArticleList", err}
	}
	return FetchTopCategoryArticletListRes{
		Articles: arts,
	}, nil
}

type FetchTopCategoryArticletListRes struct {
	Articles []data.Article
}
