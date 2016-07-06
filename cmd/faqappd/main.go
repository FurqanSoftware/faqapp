package main

import (
	"fmt"
	"log"
	"net/http"

	"git.furqan.io/faqapp/faqapp/cfg"
	"git.furqan.io/faqapp/faqapp/db"
	"github.com/gorilla/mux"
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

	r := mux.NewRouter()
	InitRouter(r, dbSess)

	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r)
	if err != nil {
		log.Fatalln("listen and serve:", err)
	}
}
