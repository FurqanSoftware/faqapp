package api

import (
	"git.furqan.io/faqapp/faqapp/db"
	"github.com/gorilla/mux"
)

func InitRouter(r *mux.Router, dbSess *db.Session) {
	r.NewRoute().
		Name("ServeAccountList").
		Methods("GET").
		Path("/accounts").
		Handler(RequireSession{
			AccountStore: db.AccountStore{Session: dbSess},
			Handler: ServeAccountList{
				AccountStore: db.AccountStore{Session: dbSess},
			},
		})
	r.NewRoute().
		Name("CreateAccount").
		Methods("POST").
		Path("/accounts").
		Handler(RequireSession{
			AccountStore: db.AccountStore{Session: dbSess},
			Handler: CreateAccount{
				AccountStore: db.AccountStore{Session: dbSess},
			},
		})

	r.NewRoute().
		Name("FetchCategoryList").
		Methods("GET").
		Path("/categories").
		Handler(RequireSession{
			AccountStore: db.AccountStore{Session: dbSess},
			Handler: ServeCategoryList{
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})
	r.NewRoute().
		Name("CreateCategory").
		Methods("POST").
		Path("/categories").
		Handler(RequireSession{
			AccountStore: db.AccountStore{Session: dbSess},
			Handler: CreateCategory{
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})

	r.NewRoute().
		Name("CreateSession").
		Methods("POST").
		Path("/sessions").
		Handler(CreateSession{
			AccountStore: db.AccountStore{Session: dbSess},
		})
}
