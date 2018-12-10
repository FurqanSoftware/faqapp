package main

import (
	"log"

	"git.furqansoftware.net/faqapp/faqapp/core"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

func CreateDefaultAccount(dbSess *db.Session) error {
	res, err := core.Do(core.CreateDefaultAccount{
		AccountStore: db.AccountStore{Session: dbSess},
	})
	if err != nil {
		return err
	}

	if res.(core.CreateDefaultAccountRes).Created {
		log.Println("default account created")
	}
	return nil
}
