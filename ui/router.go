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
		AccountStore: db.AccountStore{Session: dbSess},
		SettingStore: db.SettingStore{Session: dbSess},
		SessionStore: sessStore,
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
		Name("HandleBackLanderRedirect").
		Methods("GET").
		Path("/_").
		Handler(RequireSessionAccount{
			Handler: http.RedirectHandler("/_/categories", http.StatusSeeOther),
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
		Handler(RequireSessionAccount{
			Handler: ServeBackCategoryList{
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})
	t.NewRoute().
		Name("ServeBackCategoryNewForm").
		Methods("GET").
		Path("/_/categories/new").
		Handler(RequireSessionAccount{
			Handler: ServeBackCategoryNewForm{},
		})
	t.NewRoute().
		Name("HandleBackCategoryNewForm").
		Methods("POST").
		Path("/_/categories/new").
		Handler(RequireSessionAccount{
			Handler: HandleBackCategoryNewForm{
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})
	t.NewRoute().
		Name("ServeBackCategoryEditForm").
		Methods("GET").
		Path("/_/categories/{category_id}/edit").
		Handler(RequireSessionAccount{
			Handler: ServeBackCategoryEditForm{
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})
	t.NewRoute().
		Name("HandleBackCategoryEditForm").
		Methods("POST").
		Path("/_/categories/{category_id}/edit").
		Handler(RequireSessionAccount{
			Handler: HandleBackCategoryEditForm{
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})
	t.NewRoute().
		Name("ServeBackArticleList").
		Methods("GET").
		Path("/_/articles").
		Handler(RequireSessionAccount{
			Handler: ServeBackArticleList{
				ArticleStore:  db.ArticleStore{Session: dbSess},
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})
	t.NewRoute().
		Name("ServeBackArticleNewForm").
		Methods("GET").
		Path("/_/articles/new").
		Handler(RequireSessionAccount{
			Handler: ServeBackArticleNewForm{
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})
	t.NewRoute().
		Name("HandleBackArticleNewForm").
		Methods("POST").
		Path("/_/articles/new").
		Handler(RequireSessionAccount{
			Handler: HandleBackArticleNewForm{
				ArticleStore:  db.ArticleStore{Session: dbSess},
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})
	t.NewRoute().
		Name("ServeBackArticleEditForm").
		Methods("GET").
		Path("/_/articles/{article_id}/edit").
		Handler(RequireSessionAccount{
			Handler: ServeBackArticleEditForm{
				ArticleStore:  db.ArticleStore{Session: dbSess},
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})
	t.NewRoute().
		Name("HandleBackArticleEditForm").
		Methods("POST").
		Path("/_/articles/{article_id}/edit").
		Handler(RequireSessionAccount{
			Handler: HandleBackArticleEditForm{
				ArticleStore:  db.ArticleStore{Session: dbSess},
				CategoryStore: db.CategoryStore{Session: dbSess},
			},
		})
	t.NewRoute().
		Name("HandleBackSettingRedirect").
		Methods("GET").
		Path("/_/settings").
		Handler(RequireSessionAccount{
			Handler: http.RedirectHandler("/_/settings/password", http.StatusSeeOther),
		})
	t.NewRoute().
		Name("ServeBackSettingPasswordForm").
		Methods("GET").
		Path("/_/settings/password").
		Handler(RequireSessionAccount{
			Handler: ServeBackSettingPasswordForm{},
		})
	t.NewRoute().
		Name("HandleBackSettingPasswordForm").
		Methods("POST").
		Path("/_/settings/password").
		Handler(RequireSessionAccount{
			Handler: HandleBackSettingPasswordForm{
				AccountStore: db.AccountStore{Session: dbSess},
			},
		})
	t.NewRoute().
		Name("ServeBackSettingAdvancedForm").
		Methods("GET").
		Path("/_/settings/advanced").
		Handler(RequireSessionAccount{
			Handler: ServeBackSettingAdvancedForm{
				SettingStore: db.SettingStore{Session: dbSess},
			},
		})
	t.NewRoute().
		Name("HandleBackSettingAdvancedForm").
		Methods("POST").
		Path("/_/settings/advanced").
		Handler(RequireSessionAccount{
			Handler: HandleBackSettingAdvancedForm{
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
		Name("HandleArticleSearch").
		Methods("GET").
		Path("/search").
		Handler(HandleArticleSearch{
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
