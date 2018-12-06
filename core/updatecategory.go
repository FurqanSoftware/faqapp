package core

import (
	"gopkg.in/mgo.v2/bson"

	"git.furqansoftware.net/faqapp/faqapp/data"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

type UpdateCategory struct {
	ID    string
	Title string
	Slug  string
	Order int

	CategoryStore db.CategoryStore
}

func (a UpdateCategory) Validate() error {
	if !bson.IsObjectIdHex(a.ID) {
		return ValidationError{"UpdateCategory", "ID", IssueInvalid}
	}
	if a.Title == "" {
		return ValidationError{"UpdateCategory", "Title", IssueMissing}
	}
	if len(a.Title) > 128 {
		return ValidationError{"UpdateCategory", "Title", IssueTooLong}
	}
	if a.Slug == "" {
		return ValidationError{"UpdateCategory", "Slug", IssueMissing}
	}
	if len(a.Slug) > 128 {
		return ValidationError{"UpdateCategory", "Slug", IssueTooLong}
	}
	return nil
}

func (a UpdateCategory) Do() (Result, error) {
	cat, err := a.CategoryStore.Get(bson.ObjectIdHex(a.ID))
	if err != nil {
		return nil, DatabaseError{"UpdateCategory", err}
	}

	cat.Title = a.Title
	cat.Slug = a.Slug
	cat.Order = a.Order
	err = a.CategoryStore.Put(cat)
	if err != nil {
		return nil, DatabaseError{"UpdateCategory", err}
	}

	return UpdateCategoryRes{
		Category: cat,
	}, nil
}

type UpdateCategoryRes struct {
	Category *data.Category
}
