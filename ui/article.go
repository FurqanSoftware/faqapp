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
	if err != nil {
		log.Println("fetch category by slug:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	cat := res.(core.FetchCategoryBySlugRes).Category

	res, err = core.Do(core.FetchArticleBySlug{
		CategoryID:    cat.ID.Hex(),
		Slug:          vars["article_slug"],
		ArticleStore:  h.ArticleStore,
		CategoryStore: h.CategoryStore,
	})
	if err != nil {
		log.Println("fetch article by slug:", err)
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

var BackArticleListTpl = template.Must(template.ParseFiles("ui/gohtml/layout.gohtml", "ui/gohtml/backarticlelist.gohtml", "ui/gohtml/backtabset.gohtml"))

type ServeBackArticleList struct {
	ArticleStore  db.ArticleStore
	CategoryStore db.CategoryStore
}

func (h ServeBackArticleList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := core.Do(core.FetchArticleList{
		ArticleStore: h.ArticleStore,
	})
	if err != nil {
		log.Println("fetch article list:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	arts := res.(core.FetchArticleListRes).Articles

	artCat := map[string]*data.Category{}
	for _, art := range arts {
		res, err := core.Do(core.FetchCategory{
			ID:            art.CategoryID.Hex(),
			CategoryStore: h.CategoryStore,
		})
		if err != nil {
			log.Println("fetch category:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		artCat[art.ID.Hex()] = res.(core.FetchCategoryRes).Category
	}

	err = BackArticleListTpl.Execute(w, struct {
		Articles        []data.Article
		ArticleCategory map[string]*data.Category
	}{
		Articles:        arts,
		ArticleCategory: artCat,
	})
	if err != nil {
		log.Println("execute template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
