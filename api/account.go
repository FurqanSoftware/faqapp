package api

import (
	"log"
	"net/http"
	"time"

	"git.furqan.io/faqapp/faqapp/core"
	"git.furqan.io/faqapp/faqapp/db"
)

type ServeAccountList struct {
	AccountRepo db.Accounts
}

type ServeAccountListRes []ServeAccountListResItem

type ServeAccountListResItem struct {
	ID         string    `json:"id"`
	Handle     string    `json:"handle"`
	FirstIP    string    `json:"first_ip"`
	RecentIP   string    `json:"recent_ip"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func (h ServeAccountList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := core.Do(core.FetchAccountList{
		AccountRepo: h.AccountRepo,
	})
	if err != nil {
		log.Println("fetch account list:", err)
		HandleActionError(w, r, err)
		return
	}

	resp := ServeAccountListRes{}
	for _, acc := range res.(core.FetchAccountListRes).Accounts {
		resp = append(resp, ServeAccountListResItem{
			ID:         acc.ID.Hex(),
			Handle:     acc.Handle,
			FirstIP:    acc.FirstIP,
			RecentIP:   acc.RecentIP,
			CreatedAt:  acc.CreatedAt,
			ModifiedAt: acc.ModifiedAt,
		})
	}
	ServeResult(w, r, resp)
}

type CreateAccountVal struct {
	Handle   string `json:"handle"`
	Password string `json:"password"`
}

type CreateAccount struct {
	AccountRepo db.Accounts
}

type CreateAccountRes struct {
	ID string `json:"id"`
}

func (h CreateAccount) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := CreateAccountVal{}
	err := ParseRequestBody(r, &body)
	if err != nil {
		log.Println("parse request body:", err)
		ServeBadRequest(w, r)
		return
	}

	res, err := core.Do(core.CreateAccount{
		Handle:      body.Handle,
		Password:    body.Password,
		AccountRepo: h.AccountRepo,
	})
	if err != nil {
		log.Println("create account:", err)
		HandleActionError(w, r, err)
		return
	}

	resp := CreateAccountRes{
		ID: res.(core.CreateAccountRes).Account.ID.Hex(),
	}
	ServeResult(w, r, resp)
}
