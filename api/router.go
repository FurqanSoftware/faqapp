package api

import (
	"git.furqan.io/faqapp/faqapp/db"
	"github.com/gorilla/mux"
)

func InitRouter(r *mux.Router, dbSess *db.Session) {
	r.NewRoute().
		Name("ServeAccountList").
		Methods("GET").
		Path("/accounts").
		Handler(ServeAccountList{
			AccountRepo: db.Accounts{Session: dbSess},
		})
	r.NewRoute().
		Name("CreateAccount").
		Methods("POST").
		Path("/accounts").
		Handler(CreateAccount{
			AccountRepo: db.Accounts{Session: dbSess},
		})
}
