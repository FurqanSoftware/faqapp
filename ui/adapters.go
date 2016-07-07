package ui

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"

	"git.furqan.io/faqapp/faqapp/core"
	"git.furqan.io/faqapp/faqapp/db"
)

type RequireSession struct {
	AccountStore db.AccountStore
	SessionStore sessions.Store
	Handler      http.Handler
}

func (h RequireSession) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	claims, err := GetRequestClaims(h.SessionStore, r)
	if err != nil {
		log.Println("get request claims:", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	res, err := core.Do(core.VerifySession{
		Claims:       claims,
		AccountStore: h.AccountStore,
	})
	if err != nil {
		log.Println("verify session:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if !res.(core.VerifySessionRes).Okay {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	h.Handler.ServeHTTP(w, r)
}
