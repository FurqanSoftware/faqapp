package ui

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"git.furqan.io/faqapp/faqapp/core"
	"git.furqan.io/faqapp/faqapp/data"
	"git.furqan.io/faqapp/faqapp/db"
)

var ArticleViewTpl = template.Must(template.ParseFiles("ui/gohtml/layout.gohtml", "ui/gohtml/articleview.gohtml"))

type ServeArticleView struct {
	ArticleStore  db.ArticleStore
	CategoryStore db.CategoryStore
}

func (h ServeArticleView) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	res, err := core.Do(core.FetchCategoryBySlug{
		Slug:          vars["category_slug"],
		CategoryStore: h.CategoryStore,
	})
	cat := res.(core.FetchCategoryBySlugRes).Category

	res, err = core.Do(core.FetchArticleBySlug{
		CategoryID:    cat.ID.Hex(),
		Slug:          vars["article_slug"],
		ArticleStore:  h.ArticleStore,
		CategoryStore: h.CategoryStore,
	})
	if err != nil {
		log.Println("fetch article list:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	art := res.(core.FetchArticleBySlugRes).Article

	err = ArticleViewTpl.Execute(w, struct {
		Article  *data.Article
		Category *data.Category
	}{
		Article:  art,
		Category: cat,
	})
	if err != nil {
		log.Println("execute template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
