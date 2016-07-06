package main

import (
	"git.furqan.io/faqapp/faqapp/api"
	"git.furqan.io/faqapp/faqapp/db"
	"github.com/gorilla/mux"
)

func InitRouter(r *mux.Router, dbSess *db.Session) {
	apiR := r.NewRoute().
		PathPrefix("/api").
		Subrouter()
	api.InitRouter(apiR, dbSess)
}
