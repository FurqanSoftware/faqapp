package core

import "git.furqan.io/faqapp/faqapp/db"

type CreateAccount struct {
	Handle   string
	Password string

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

func (a CreateAccount) Do() (Result, error) {
	return nil, nil
}
