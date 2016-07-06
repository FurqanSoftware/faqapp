package core

import (
	"git.furqan.io/faqapp/faqapp/data"
	"git.furqan.io/faqapp/faqapp/db"
)

type FetchArticleBySlug struct {
	Slug string

	ArticleStore db.ArticleStore
}

func (a FetchArticleBySlug) Do() (Result, error) {
	art, err := a.ArticleStore.GetBySlug(a.Slug)
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
