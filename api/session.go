package api

import (
	"log"
	"net/http"

	"git.furqan.io/faqapp/faqapp/core"
	"git.furqan.io/faqapp/faqapp/db"
)

type CreateSession struct {
	AccountRepo db.Accounts
}

type CreateSessionVal struct {
	Handle   string `json:"handle"`
	Password string `json:"password"`
}

type CreateSessionRes struct {
	Match bool   `json:"match"`
	Token string `json:"token"`
}

func (h CreateSession) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := CreateSessionVal{}
	err := ParseRequestBody(r, &body)
	if err != nil {
		log.Println("parse request body:", err)
		ServeBadRequest(w, r)
		return
	}

	res, err := core.Do(core.CreateSession{
		Handle:      body.Handle,
		Password:    body.Password,
		AccountRepo: h.AccountRepo,
	})
	if err != nil {
		log.Println("create session:", err)
		HandleActionError(w, r, err)
		return
	}

	csRes := res.(core.CreateSessionRes)
	resp := CreateSessionRes{
		Match: csRes.Match,
		Token: csRes.Token,
	}
	ServeResult(w, r, resp)
}
