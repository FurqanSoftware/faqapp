package ui

import (
	"net/http"

	"git.furqan.io/faqapp/faqapp/db"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func InitRouter(r *mux.Router, dbSess *db.Session, sessStore sessions.Store) {
	t := mux.NewRouter()
	r.NewRoute().Handler(PrepareContext{
		SettingStore: db.SettingStore{Session: dbSess},
		Handler:      t,
	})

	t.NewRoute().
		Name("ServeCustomCSS").
		Methods("GET").
		Path("/custom.css").
		Handler(ServeCustomCSS{
			SettingStore: db.SettingStore{Session: dbSess},
		})

	t.NewRoute().
		Name("ServeLoginForm").
		Methods("GET").
		Path("/_").
		Handler(RequireSession{
			AccountStore: db.AccountStore{Session: dbSess},
			SessionStore: sessStore,
			Handler:      http.RedirectHandler("/_/categories", http.StatusSeeOther),
		})
	t.NewRoute().
		Name("ServeLoginForm").
		Methods("GET").
		Path("/_/login").
		Handler(ServeLoginForm{})
	t.NewRoute().
		Name("HandleLoginForm").
		Methods("POST").
		Path("/_/login").
		Handler(HandleLoginForm{
			AccountStore: db.AccountStore{Session: dbSess},
			SessionStore: sessStore,
		})
	t.NewRoute().
		Name("HandleLoginForm").
		Methods("GET").
		Path("/_/logout").
		Handler(HandleLogout{
			SessionStore: sessStore,
		})
	t.NewRoute().
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
	t.NewRoute().
		Name("ServeBackCategoryNewForm").
		Methods("GET").
		Path("/_/categories/new").
		Handler(RequireSession{
			AccountStore: db.AccountStore{Session: dbSess},
			SessionStore: sessStore,
			Handler:      ServeBackCategoryNewForm{},
		})
	t.NewRoute().
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
	t.NewRoute().
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
	t.NewRoute().
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
	t.NewRoute().
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
	t.NewRoute().
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
	t.NewRoute().
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
	t.NewRoute().
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
	t.NewRoute().
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
	t.NewRoute().
		Name("ServeBackSettingForm").
		Methods("GET").
		Path("/_/settings").
		Handler(RequireSession{
			AccountStore: db.AccountStore{Session: dbSess},
			SessionStore: sessStore,
			Handler: ServeBackSettingForm{
				SettingStore: db.SettingStore{Session: dbSess},
			},
		})
	t.NewRoute().
		Name("HandleBackSettingForm").
		Methods("POST").
		Path("/_/settings").
		Handler(RequireSession{
			AccountStore: db.AccountStore{Session: dbSess},
			SessionStore: sessStore,
			Handler: HandleBackSettingForm{
				SettingStore: db.SettingStore{Session: dbSess},
			},
		})

	t.NewRoute().
		Name("ServeAccountList").
		Methods("GET").
		Path("/").
		Handler(ServeHomepage{
			ArticleStore:  db.ArticleStore{Session: dbSess},
			CategoryStore: db.CategoryStore{Session: dbSess},
		})
	t.NewRoute().
		Name("ServeCategoryView").
		Methods("GET").
		Path("/{category_slug}").
		Handler(ServeCategoryView{
			ArticleStore:  db.ArticleStore{Session: dbSess},
			CategoryStore: db.CategoryStore{Session: dbSess},
		})
	t.NewRoute().
		Name("ServeArticleView").
		Methods("GET").
		Path("/{category_slug}/{article_slug}").
		Handler(ServeArticleView{
			ArticleStore:  db.ArticleStore{Session: dbSess},
			CategoryStore: db.CategoryStore{Session: dbSess},
		})
}
