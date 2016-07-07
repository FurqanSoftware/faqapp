package main

import (
	"git.furqan.io/faqapp/faqapp/api"
	"git.furqan.io/faqapp/faqapp/db"
	"git.furqan.io/faqapp/faqapp/ui"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func InitRouter(r *mux.Router, dbSess *db.Session, sessStore sessions.Store) {
	apiR := r.NewRoute().
		PathPrefix("/api").
		Subrouter()
	api.InitRouter(apiR, dbSess)

	ui.InitRouter(r, dbSess, sessStore)
}
