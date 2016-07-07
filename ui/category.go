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

var CategoryViewTpl = template.Must(template.ParseFiles("ui/gohtml/layout.gohtml", "ui/gohtml/categoryview.gohtml"))

type ServeCategoryView struct {
	ArticleStore  db.ArticleStore
	CategoryStore db.CategoryStore
}

func (h ServeCategoryView) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	res, err = core.Do(core.FetchArticleList{
		CategoryID:    cat.ID.Hex(),
		CategoryStore: h.CategoryStore,
		ArticleStore:  h.ArticleStore,
	})
	if err != nil {
		log.Println("fetch article list:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	arts := res.(core.FetchArticleListRes).Articles

	err = CategoryViewTpl.Execute(w, struct {
		Category *data.Category
		Articles []data.Article
	}{
		Category: cat,
		Articles: arts,
	})
	if err != nil {
		log.Println("execute template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

var BackCategoryListTpl = template.Must(template.ParseFiles("ui/gohtml/layout.gohtml", "ui/gohtml/backcategorylist.gohtml", "ui/gohtml/backtabset.gohtml"))

type ServeBackCategoryList struct {
	CategoryStore db.CategoryStore
}

func (h ServeBackCategoryList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := core.Do(core.FetchCategoryList{
		CategoryStore: h.CategoryStore,
	})
	if err != nil {
		log.Println("fetch category list:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	cats := res.(core.FetchCategoryListRes).Categories

	err = BackCategoryListTpl.Execute(w, struct {
		Categories []data.Category
	}{
		Categories: cats,
	})
	if err != nil {
		log.Println("execute template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

var BackCategoryNewFormTpl = template.Must(template.ParseFiles("ui/gohtml/layout.gohtml", "ui/gohtml/backcategorynewform.gohtml", "ui/gohtml/backtabset.gohtml"))

type ServeBackCategoryNewForm struct {
}

func (h ServeBackCategoryNewForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := BackCategoryNewFormTpl.Execute(w, struct{}{})
	if err != nil {
		log.Println("execute template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

type HandleBackCategoryNewForm struct {
	CategoryStore db.CategoryStore
}

type HandleBackCategoryNewFormVal struct {
	Title string `schema:"title"`
	Slug  string `schema:"slug"`
	Order int    `schema:"order"`
}

func (h HandleBackCategoryNewForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := HandleBackCategoryNewFormVal{}
	err := ParseRequestBody(r, &body)
	if err != nil {
		log.Println("parse request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	_, err = core.Do(core.CreateCategory{
		Title:         body.Title,
		Slug:          body.Slug,
		Order:         body.Order,
		CategoryStore: h.CategoryStore,
	})
	if err != nil {
		log.Println("create session:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/_/categories", http.StatusSeeOther)
}
