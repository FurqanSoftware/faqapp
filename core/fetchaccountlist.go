package core

import "git.furqan.io/faqapp/faqapp/db"

type FetchAccountList struct {
	AccountRepo db.Accounts
}

func (a FetchAccountList) Do() (Result, error) {
	return nil, nil
}
