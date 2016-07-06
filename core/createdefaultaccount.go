package core

import (
	"git.furqan.io/faqapp/faqapp/data"
	"git.furqan.io/faqapp/faqapp/db"
)

type CreateDefaultAccount struct {
	AccountStore db.AccountStore
}

func (a CreateDefaultAccount) Do() (res Result, err error) {
	accs, err := a.AccountStore.List(0, 1)
	if err != nil {
		return nil, DatabaseError{"CreateDefaultAccount", err}
	}
	if len(accs) != 0 {
		return CreateDefaultAccountRes{
			Created: false,
		}, nil
	}

	acc := data.Account{
		Handle: "faqapp",
	}
	acc.Password, err = data.NewAccountPassword("p@ssword")
	if err != nil {
		return nil, err
	}
	err = a.AccountStore.Put(&acc)
	if err != nil {
		return nil, DatabaseError{"CreateDefaultAccount", err}
	}

	return CreateDefaultAccountRes{
		Created: true,
	}, nil
}

type CreateDefaultAccountRes struct {
	Created bool
}
