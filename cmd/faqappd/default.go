package main

import (
	"log"

	"git.furqan.io/faqapp/faqapp/core"
	"git.furqan.io/faqapp/faqapp/db"
)

func CreateDefaultAccount(dbSess *db.Session) error {
	res, err := core.Do(core.CreateDefaultAccount{
		AccountRepo: db.Accounts{Session: dbSess},
	})
	if res.(core.CreateDefaultAccountRes).Created {
		log.Println("default account created")
	}
	return err
}
