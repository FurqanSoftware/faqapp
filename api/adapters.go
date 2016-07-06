package api

import (
	"log"
	"net/http"

	"git.furqan.io/faqapp/faqapp/core"
	"git.furqan.io/faqapp/faqapp/db"
)

type RequireSession struct {
	AccountStore db.AccountStore
	Handler      http.Handler
}

func (h RequireSession) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	claims, err := GetRequestClaims(r)
	if err != nil {
		log.Println("get request claims:", err)
		ServeUnauthorized(w, r)
		return
	}
	res, err := core.Do(core.VerifySession{
		Claims:       claims,
		AccountStore: h.AccountStore,
	})
	if err != nil {
		log.Println("verify session:", err)
		HandleActionError(w, r, err)
		return
	}
	if !res.(core.VerifySessionRes).Okay {
		ServeUnauthorized(w, r)
		return
	}

	h.Handler.ServeHTTP(w, r)
}
