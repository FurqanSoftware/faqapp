package core

import (
	"math"

	"git.furqan.io/faqapp/faqapp/data"
	"git.furqan.io/faqapp/faqapp/db"
)

type FetchCategoryList struct {
	CategoryStore db.CategoryStore
}

func (a FetchCategoryList) Do() (Result, error) {
	cats, err := a.CategoryStore.List(0, math.MaxInt32)
	if err != nil {
		return nil, DatabaseError{"FetchCategoryList", err}
	}
	return FetchCategoryListRes{
		Categories: cats,
	}, nil
}

type FetchCategoryListRes struct {
	Categories []data.Category
}
