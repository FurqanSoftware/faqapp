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
		Name("HandleLoginForm").
		Methods("GET").
		Path("/_/logout").
		Handler(HandleLogout{
			SessionStore: sessStore,
		})
	r.NewRoute().
		Name("ServeBackCategoryList").
		Methods("GET").
		Path("/_/categories").
		Handler(RequireSession{
			AccountStore: db.AccountStore{Session: dbSess},
			SessionStore: sessStore,
			Handler: ServeBackCategoryList{
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})
	r.NewRoute().
		Name("ServeBackCategoryNewForm").
		Methods("GET").
		Path("/_/categories/new").
		Handler(RequireSession{
			AccountStore: db.AccountStore{Session: dbSess},
			SessionStore: sessStore,
			Handler:      ServeBackCategoryNewForm{},
		})
	r.NewRoute().
		Name("HandleBackCategoryNewForm").
		Methods("POST").
		Path("/_/categories/new").
		Handler(RequireSession{
			AccountStore: db.AccountStore{Session: dbSess},
			SessionStore: sessStore,
			Handler: HandleBackCategoryNewForm{
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})
	r.NewRoute().
		Name("ServeBackCategoryEditForm").
		Methods("GET").
		Path("/_/categories/{category_id}/edit").
		Handler(RequireSession{
			AccountStore: db.AccountStore{Session: dbSess},
			SessionStore: sessStore,
			Handler: ServeBackCategoryEditForm{
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})
	r.NewRoute().
		Name("HandleBackCategoryEditForm").
		Methods("POST").
		Path("/_/categories/{category_id}/edit").
		Handler(RequireSession{
			AccountStore: db.AccountStore{Session: dbSess},
			SessionStore: sessStore,
			Handler: HandleBackCategoryEditForm{
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})
	r.NewRoute().
		Name("ServeBackArticleList").
		Methods("GET").
		Path("/_/articles").
		Handler(RequireSession{
			AccountStore: db.AccountStore{Session: dbSess},
			SessionStore: sessStore,
			Handler: ServeBackArticleList{
				ArticleStore:  db.ArticleStore{Session: dbSess},
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})
	r.NewRoute().
		Name("ServeBackArticleNewForm").
		Methods("GET").
		Path("/_/articles/new").
		Handler(RequireSession{
			AccountStore: db.AccountStore{Session: dbSess},
			SessionStore: sessStore,
			Handler: ServeBackArticleNewForm{
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})
	r.NewRoute().
		Name("HandleBackArticleNewForm").
		Methods("POST").
		Path("/_/articles/new").
		Handler(RequireSession{
			AccountStore: db.AccountStore{Session: dbSess},
			SessionStore: sessStore,
			Handler: HandleBackArticleNewForm{
				ArticleStore:  db.ArticleStore{Session: dbSess},
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})
	r.NewRoute().
		Name("ServeBackArticleEditForm").
		Methods("GET").
		Path("/_/articles/{article_id}/edit").
		Handler(RequireSession{
			AccountStore: db.AccountStore{Session: dbSess},
			SessionStore: sessStore,
			Handler: ServeBackArticleEditForm{
				ArticleStore:  db.ArticleStore{Session: dbSess},
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})
	r.NewRoute().
		Name("HandleBackArticleEditForm").
		Methods("POST").
		Path("/_/articles/{article_id}/edit").
		Handler(RequireSession{
			AccountStore: db.AccountStore{Session: dbSess},
			SessionStore: sessStore,
			Handler: HandleBackArticleEditForm{
				ArticleStore:  db.ArticleStore{Session: dbSess},
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
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
