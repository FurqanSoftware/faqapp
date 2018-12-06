package ui

import (
	"html/template"
	"log"
	"net/http"

	"git.furqansoftware.net/faqapp/faqapp/core"
	"git.furqansoftware.net/faqapp/faqapp/data"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

var HomepageTpl = template.Must(template.ParseFiles("ui/gohtml/layout.gohtml", "ui/gohtml/homepage.gohtml"))

type ServeHomepage struct {
	ArticleStore  db.ArticleStore
	CategoryStore db.CategoryStore
}

func (h ServeHomepage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := GetContext(r)

	res, err := core.Do(core.FetchCategoryList{
		CategoryStore: h.CategoryStore,
	})
	if err != nil {
		log.Println("fetch category list:", err)
		HandleActionError(w, r, err)
		return
	}
	cats := res.(core.FetchCategoryListRes).Categories

	topCatArts := map[string][]data.Article{}
	for _, cat := range cats {
		res, err := core.Do(core.FetchTopCategoryArticleList{
			CategoryID:    cat.ID.Hex(),
			ArticleStore:  h.ArticleStore,
			CategoryStore: h.CategoryStore,
		})
		if err != nil {
			log.Println("fetch top category article list:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		topCatArts[cat.ID.Hex()] = res.(core.FetchTopCategoryArticleListRes).Articles
	}

	err = ExecuteTemplate(HomepageTpl, w, struct {
		Context             Context
		Categories          []data.Category
		TopCategoryArticles map[string][]data.Article
	}{
		Context:             ctx,
		Categories:          cats,
		TopCategoryArticles: topCatArts,
	})
	if err != nil {
		log.Println("execute template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
