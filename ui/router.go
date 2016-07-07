package ui

import (
	"git.furqan.io/faqapp/faqapp/db"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func InitRouter(r *mux.Router, dbSess *db.Session, sessStore sessions.Store) {

	r.NewRoute().
		Name("ServeLoginForm").
		Methods("GET").
		Path("/_/login").
		Handler(ServeLoginForm{})
	r.NewRoute().
		Name("HandleLoginForm").
		Methods("POST").
		Path("/_/login").
		Handler(HandleLoginForm{
			AccountStore: db.AccountStore{Session: dbSess},
			SessionStore: sessStore,
		})

	r.NewRoute().
		Name("ServeAccountList").
		Methods("GET").
		Path("/").
		Handler(ServeHomepage{
			ArticleStore:  db.ArticleStore{Session: dbSess},
			CategoryStore: db.CategoryStore{Session: dbSess},
		})
	r.NewRoute().
		Name("ServeCategoryView").
		Methods("GET").
		Path("/{category_slug}").
		Handler(ServeCategoryView{
			ArticleStore:  db.ArticleStore{Session: dbSess},
			CategoryStore: db.CategoryStore{Session: dbSess},
		})
	r.NewRoute().
		Name("ServeArticleView").
		Methods("GET").
		Path("/{category_slug}/{article_slug}").
		Handler(ServeArticleView{
			ArticleStore:  db.ArticleStore{Session: dbSess},
			CategoryStore: db.CategoryStore{Session: dbSess},
		})
}
