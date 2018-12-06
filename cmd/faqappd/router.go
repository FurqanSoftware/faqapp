package main

import (
	"git.furqansoftware.net/faqapp/faqapp/api"
	"git.furqansoftware.net/faqapp/faqapp/db"
	"git.furqansoftware.net/faqapp/faqapp/ui"
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
