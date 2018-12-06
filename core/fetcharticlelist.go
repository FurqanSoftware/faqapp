package core

import (
	"math"

	"git.furqansoftware.net/faqapp/faqapp/data"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

type FetchArticleList struct {
	ArticleStore db.ArticleStore
}

func (a FetchArticleList) Do() (Result, error) {
	arts, err := a.ArticleStore.List(0, math.MaxInt32)
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
