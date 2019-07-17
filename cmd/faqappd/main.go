package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"git.furqansoftware.net/faqapp/faqapp/cfg"
	"git.furqansoftware.net/faqapp/faqapp/core"
	"git.furqansoftware.net/faqapp/faqapp/db"
	"git.furqansoftware.net/faqapp/faqapp/health"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func main() {
	err := cfg.Load()
	if err != nil {
		log.Fatalln("load configuration:", err)
	}

	dbSess, err := db.Open(cfg.MongoURL)
	if err != nil {
		log.Fatalln("open database session:", err)
	}
	err = db.MakeIndexes(dbSess)
	if err != nil {
		log.Fatalln("make indexes:", err)
	}

	_, err = core.Do(core.PrepopulateSettings{
		SettingStore: db.SettingStore{Session: dbSess},
	})
	if err != nil {
		log.Println("prepopulate settings:", err)
		return
	}

	err = CreateDefaultAccount(dbSess)
	if err != nil {
		log.Fatalln("create default account:", err)
	}

	sessStore := sessions.NewCookieStore([]byte(cfg.Secret))

	go health.Loop(11*time.Second, dbSess)

	r := mux.NewRouter()
	InitRouter(r, dbSess, sessStore)

	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r)
	if err != nil {
		log.Fatalln("listen and serve:", err)
	}
}
