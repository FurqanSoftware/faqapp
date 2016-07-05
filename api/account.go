package api

import (
	"log"
	"net/http"

	"git.furqan.io/faqapp/faqapp/core"
)

type CreateAccountVal struct {
	Handle   string `json:"handle"`
	Password string `json:"password"`
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	body := CreateAccountVal{}
	err := ParseRequestBody(r, &body)
	if err != nil {
		log.Print("ParseRequestBody:", err)
		ServeBadRequest(w, r)
	}

	res, err := core.Do(core.CreateAccount{
		Handle:   body.Handle,
		Password: body.Password,
	})
	if err != nil {
		log.Print("CreateAccount:", err)
		ServeInternalServerError(w, r)
	}

	ServeResult(w, r, res)
}
