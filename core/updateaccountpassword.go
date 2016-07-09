package core

import (
	"gopkg.in/mgo.v2/bson"

	"git.furqan.io/faqapp/faqapp/data"
	"git.furqan.io/faqapp/faqapp/db"
)

type UpdateAccountPassword struct {
	ID      string
	Current string
	New     string
	Confirm string

	AccountStore db.AccountStore
}

func (a UpdateAccountPassword) Validate() error {
	if !bson.IsObjectIdHex(a.ID) {
		return ValidationError{"UpdateAccountPassword", "ID", IssueInvalid}
	}
	if len(a.New) < 8 {
		return ValidationError{"UpdateAccountPassword", "New", IssueTooShort}
	}
	if a.New != a.Confirm {
		return ValidationError{"UpdateAccountPassword", "Confirm", IssueInvalid}
	}
	return nil
}

func (a UpdateAccountPassword) Do() (Result, error) {
	acc, err := a.AccountStore.Get(bson.ObjectIdHex(a.ID))
	if err != nil {
		return nil, DatabaseError{"UpdateAccountPassword", err}
	}

	if !acc.Password.Match(a.Current) {
		return nil, ValidationError{"UpdateAccountPassword", "Current", IssueInvalid}
	}

	acc.Password, err = data.NewAccountPassword(a.New)
	err = a.AccountStore.Put(acc)
	if err != nil {
		return nil, DatabaseError{"UpdateAccountPassword", err}
	}

	return UpdateAccountPasswordRes{
		Account: acc,
	}, nil
}

type UpdateAccountPasswordRes struct {
	Account *data.Account
}
