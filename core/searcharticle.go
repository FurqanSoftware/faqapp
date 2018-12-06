package core

import (
	"math"

	"git.furqansoftware.net/faqapp/faqapp/data"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

type SearchArticle struct {
	Query string

	ArticleStore db.ArticleStore
}

func (a SearchArticle) Do() (Result, error) {
	arts, err := a.ArticleStore.Search(a.Query, 0, math.MaxInt32)
	if err != nil {
		return nil, DatabaseError{"SearchArticle", err}
	}
	return SearchArticleRes{
		Articles: arts,
	}, nil
}

type SearchArticleRes struct {
	Articles []data.Article
}
