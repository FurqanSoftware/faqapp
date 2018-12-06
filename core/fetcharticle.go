package core

import (
	"gopkg.in/mgo.v2/bson"

	"git.furqansoftware.net/faqapp/faqapp/data"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

type FetchArticle struct {
	ID string

	ArticleStore db.ArticleStore
}

func (a FetchArticle) Validate() error {
	if !bson.IsObjectIdHex(a.ID) {
		return ValidationError{"FetchArticle", "ID", IssueInvalid}
	}
	return nil
}

func (a FetchArticle) Do() (Result, error) {
	art, err := a.ArticleStore.Get(bson.ObjectIdHex(a.ID))
	if err != nil {
		return nil, DatabaseError{"FetchArticle", err}
	}
	return FetchArticleRes{
		Article: art,
	}, nil
}

type FetchArticleRes struct {
	Article *data.Article
}
