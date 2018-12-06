package core

import (
	"git.furqansoftware.net/faqapp/faqapp/data"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

type CreateCategory struct {
	Title string
	Slug  string
	Order int

	CategoryStore db.CategoryStore
}

func (a CreateCategory) Validate() error {
	if a.Title == "" {
		return ValidationError{"CreateCategory", "Title", IssueMissing}
	}
	if len(a.Title) > 128 {
		return ValidationError{"CreateCategory", "Title", IssueTooLong}
	}
	if a.Slug == "" {
		return ValidationError{"CreateCategory", "Slug", IssueMissing}
	}
	if len(a.Slug) > 128 {
		return ValidationError{"CreateCategory", "Slug", IssueTooLong}
	}
	return nil
}

func (a CreateCategory) Do() (res Result, err error) {
	cat := data.Category{
		Title: a.Title,
		Slug:  a.Slug,
		Order: a.Order,
	}
	err = a.CategoryStore.Put(&cat)
	if err != nil {
		return nil, DatabaseError{"CreateCategory", err}
	}
	return CreateCategoryRes{
		Category: &cat,
	}, nil
}

type CreateCategoryRes struct {
	Category *data.Category
}
