package ui

import (
	"git.furqan.io/faqapp/faqapp/db"
	"github.com/gorilla/mux"
)

func InitRouter(r *mux.Router, dbSess *db.Session) {
	r.NewRoute().
		Name("ServeAccountList").
		Methods("GET").
		Path("/").
		Handler(ServeHomepage{
			CategoryStore: db.CategoryStore{Session: dbSess},
		})
}
