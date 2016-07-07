package main

import (
	"fmt"
	"log"
	"net/http"

	"git.furqan.io/faqapp/faqapp/cfg"
	"git.furqan.io/faqapp/faqapp/db"
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

	err = CreateDefaultAccount(dbSess)
	if err != nil {
		log.Fatalln("create default account:", err)
	}

	sessStore := sessions.NewCookieStore([]byte(cfg.Secret))

	r := mux.NewRouter()
	InitRouter(r, dbSess, sessStore)

	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r)
	if err != nil {
		log.Fatalln("listen and serve:", err)
	}
}
