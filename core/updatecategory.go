package core

import (
	"gopkg.in/mgo.v2/bson"

	"git.furqan.io/faqapp/faqapp/data"
	"git.furqan.io/faqapp/faqapp/db"
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
