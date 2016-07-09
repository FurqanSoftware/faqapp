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
	SettingStore db.SettingStore
	Handler      http.Handler
}

func (h PrepareContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := Context{}

	res, err := core.Do(core.FetchSettingList{
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
		HandleActionError(w, r, err)
		return
	}
	if !res.(core.VerifySessionRes).Okay {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	h.Handler.ServeHTTP(w, r)
}
