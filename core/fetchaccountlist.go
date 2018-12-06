package core

import (
	"math"

	"git.furqansoftware.net/faqapp/faqapp/data"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

type FetchAccountList struct {
	AccountStore db.AccountStore
}

func (a FetchAccountList) Do() (Result, error) {
	accs, err := a.AccountStore.List(0, math.MaxInt32)
	if err != nil {
		return nil, DatabaseError{"FetchAccountList", err}
	}
	return FetchAccountListRes{
		Accounts: accs,
	}, nil
}

type FetchAccountListRes struct {
	Accounts []data.Account
}
