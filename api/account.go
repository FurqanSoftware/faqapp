package api

import (
	"log"
	"net/http"

	"git.furqan.io/faqapp/faqapp/core"
	"git.furqan.io/faqapp/faqapp/db"
)

type ServeAccountList struct {
	AccountRepo db.Accounts
}

func (h ServeAccountList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := core.Do(core.FetchAccountList{})
	if err != nil {
		log.Print("fetch account list:", err)
		ServeInternalServerError(w, r)
	}

	ServeResult(w, r, res)
}

type CreateAccountVal struct {
	Handle   string `json:"handle"`
	Password string `json:"password"`
}

type CreateAccount struct {
	AccountRepo db.Accounts
}

func (h CreateAccount) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := CreateAccountVal{}
	err := ParseRequestBody(r, &body)
	if err != nil {
		log.Print("parse request body:", err)
		ServeBadRequest(w, r)
	}

	res, err := core.Do(core.CreateAccount{
		Handle:   body.Handle,
		Password: body.Password,
	})
	if err != nil {
		log.Print("create account:", err)
		ServeInternalServerError(w, r)
	}

	ServeResult(w, r, res)
}
