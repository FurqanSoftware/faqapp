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
			ArticleStore:  db.ArticleStore{Session: dbSess},
			CategoryStore: db.CategoryStore{Session: dbSess},
		})
}
