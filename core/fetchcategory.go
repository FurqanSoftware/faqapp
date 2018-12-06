package core

import (
	"gopkg.in/mgo.v2/bson"

	"git.furqansoftware.net/faqapp/faqapp/data"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

type FetchCategory struct {
	ID string

	CategoryStore db.CategoryStore
}

func (a FetchCategory) Validate() error {
	if !bson.IsObjectIdHex(a.ID) {
		return ValidationError{"FetchCategory", "ID", IssueInvalid}
	}
	return nil
}

func (a FetchCategory) Do() (Result, error) {
	cat, err := a.CategoryStore.Get(bson.ObjectIdHex(a.ID))
	if err != nil {
		return nil, DatabaseError{"FetchCategory", err}
	}
	return FetchCategoryRes{
		Category: cat,
	}, nil
}

type FetchCategoryRes struct {
	Category *data.Category
}
