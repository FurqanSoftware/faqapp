package core

import (
	"git.furqan.io/faqapp/faqapp/data"
	"git.furqan.io/faqapp/faqapp/db"
)

type CreateAccount struct {
	Handle   string
	Password string
	FirstIP  string

	AccountRepo db.Accounts
}

func (a CreateAccount) Validate() error {
	if a.Handle == "" {
		return ValidationError{"CreateAccount", "Handle", IssueMissing}
	}
	if len(a.Handle) > 14 {
		return ValidationError{"CreateAccount", "Handle", IssueTooLong}
	}
	if len(a.Password) < 8 {
		return ValidationError{"CreateAccount", "Password", IssueTooShort}
	}
	return nil
}

func (a CreateAccount) Do() (res Result, err error) {
	acc := data.Account{
		Handle:   a.Handle,
		FirstIP:  a.FirstIP,
		RecentIP: a.FirstIP,
	}
	acc.Password, err = data.NewAccountPassword(a.Password)
	if err != nil {
		return nil, err
	}
	err = a.AccountRepo.Put(&acc)
	if err != nil {
		return nil, DatabaseError{"CreateAccount", err}
	}
	return CreateAccountRes{
		Account: &acc,
	}, nil
}

type CreateAccountRes struct {
	Account *data.Account
}
