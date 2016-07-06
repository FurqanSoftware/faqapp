package ui

import (
	"html/template"
	"log"
	"net/http"

	"git.furqan.io/faqapp/faqapp/core"
	"git.furqan.io/faqapp/faqapp/data"
	"git.furqan.io/faqapp/faqapp/db"
)

var HomepageTpl = template.Must(template.ParseFiles("ui/gohtml/layout.gohtml", "ui/gohtml/homepage.gohtml"))

type ServeHomepage struct {
	CategoryStore db.CategoryStore
}

func (h ServeHomepage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := core.Do(core.FetchCategoryList{
		CategoryStore: h.CategoryStore,
	})
	if err != nil {
		log.Println("fetch category list:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	cats := res.(core.FetchCategoryListRes).Categories

	HomepageTpl.Execute(w, struct {
		Categories []data.Category
	}{
		Categories: cats,
	})
}
