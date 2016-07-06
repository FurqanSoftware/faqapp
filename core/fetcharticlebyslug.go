package core

import (
	"git.furqan.io/faqapp/faqapp/data"
	"git.furqan.io/faqapp/faqapp/db"
	"gopkg.in/mgo.v2/bson"
)

type FetchArticleBySlug struct {
	CategoryID string
	Slug       string

	ArticleStore  db.ArticleStore
	CategoryStore db.CategoryStore
}

func (a FetchArticleBySlug) Do() (Result, error) {
	cat, err := a.CategoryStore.Get(bson.ObjectIdHex(a.CategoryID))
	if err != nil {
		return nil, DatabaseError{"CreateArticle", err}
	}

	art, err := a.ArticleStore.GetBySlug(cat.ID, a.Slug)
	if err != nil {
		return nil, DatabaseError{"FetchArticleBySlug", err}
	}
	return FetchArticleBySlugRes{
		Article: art,
	}, nil
}

type FetchArticleBySlugRes struct {
	Article *data.Article
}
