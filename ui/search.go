package ui

import (
	"html/template"
	"log"
	"net/http"

	"git.furqansoftware.net/faqapp/faqapp/core"
	"git.furqansoftware.net/faqapp/faqapp/data"
	"git.furqansoftware.net/faqapp/faqapp/db"
	"github.com/gorilla/sessions"
)

var ArticleSearchTpl = template.Must(template.ParseFiles("ui/gohtml/layout.gohtml", "ui/gohtml/searchview.gohtml"))

type HandleArticleSearch struct {
	ArticleStore  db.ArticleStore
	CategoryStore db.CategoryStore
	SessionStore  sessions.Store
}

type HandleArticleSearchVal struct {
	Query string `schema:"q"`
}

func (h HandleArticleSearch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := GetContext(r)

	body := HandleArticleSearchVal{}
	err := ParseRequestBody(r, &body)
	if err != nil {
		log.Println("parse request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	res, err := core.Do(core.SearchArticle{
		Query:        body.Query,
		ArticleStore: h.ArticleStore,
	})
	if err != nil {
		log.Println("search article:", err)
		HandleActionError(w, r, err)
		return
	}
	arts := res.(core.SearchArticleRes).Articles

	artCat := map[string]*data.Category{}
	for _, art := range arts {
		res, err := core.Do(core.FetchCategory{
			ID:            art.CategoryID.Hex(),
			CategoryStore: h.CategoryStore,
		})
		if err != nil {
			log.Println("fetch category:", err)
			HandleActionError(w, r, err)
			return
		}
		artCat[art.ID.Hex()] = res.(core.FetchCategoryRes).Category
	}

	err = ExecuteTemplate(ArticleSearchTpl, w, struct {
		Page
		Context         Context
		Query           string
		Articles        []data.Article
		ArticleCategory map[string]*data.Category
	}{
		Page: Page{
			Title: body.Query + " | Articles" + " | Toph Help",
		},
		Context:         ctx,
		Query:           body.Query,
		Articles:        arts,
		ArticleCategory: artCat,
	})
	if err != nil {
		log.Println("execute template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
