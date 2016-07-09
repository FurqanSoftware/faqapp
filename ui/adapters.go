package ui

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"

	"git.furqan.io/faqapp/faqapp/core"
	"git.furqan.io/faqapp/faqapp/db"
)

type PrepareContext struct {
	AccountStore db.AccountStore
	SettingStore db.SettingStore
	SessionStore sessions.Store
	Handler      http.Handler
}

func (h PrepareContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := Context{}

	claims, err := GetRequestClaims(h.SessionStore, r)
	if err != nil {
		log.Println("get request claims:", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	res, err := core.Do(core.FetchSessionAccount{
		Claims:       claims,
		AccountStore: h.AccountStore,
	})
	if err != nil {
		log.Println("verify session:", err)
		HandleActionError(w, r, err)
		return
	}
	fsaRes := res.(core.FetchSessionAccountRes)
	if fsaRes.Account != nil {
		ctx.Account = fsaRes.Account
	}

	res, err = core.Do(core.FetchSettingList{
		SettingStore: h.SettingStore,
	})
	if err != nil {
		log.Println("fetch setting list:", err)
		HandleActionError(w, r, err)
		return
	}

	ctx.Settings = map[string]interface{}{}
	for _, stt := range res.(core.FetchSettingListRes).Settings {
		ctx.Settings[stt.Key] = stt.Value
	}

	context.Set(r, "context", ctx)

	h.Handler.ServeHTTP(w, r)
}

type RequireSessionAccount struct {
	Handler http.Handler
}

func (h RequireSessionAccount) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := GetContext(r)
	if ctx.Account == nil {
		http.Redirect(w, r, "/_/login", http.StatusSeeOther)
		return
	}

	h.Handler.ServeHTTP(w, r)
}
