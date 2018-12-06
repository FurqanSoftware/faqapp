package core

import (
	"git.furqansoftware.net/faqapp/faqapp/data"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

type FetchCategoryBySlug struct {
	Slug string

	CategoryStore db.CategoryStore
}

func (a FetchCategoryBySlug) Do() (Result, error) {
	cat, err := a.CategoryStore.GetBySlug(a.Slug)
	if err != nil {
		return nil, DatabaseError{"FetchCategoryBySlug", err}
	}
	return FetchCategoryBySlugRes{
		Category: cat,
	}, nil
}

type FetchCategoryBySlugRes struct {
	Category *data.Category
}
