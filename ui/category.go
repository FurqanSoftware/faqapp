package ui

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"git.furqansoftware.net/faqapp/faqapp/core"
	"git.furqansoftware.net/faqapp/faqapp/data"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

var CategoryViewTpl = template.Must(template.ParseFiles("ui/gohtml/layout.gohtml", "ui/gohtml/categoryview.gohtml"))

type ServeCategoryView struct {
	ArticleStore  db.ArticleStore
	CategoryStore db.CategoryStore
}

func (h ServeCategoryView) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := GetContext(r)

	vars := mux.Vars(r)

	res, err := core.Do(core.FetchCategoryBySlug{
		Slug:          vars["category_slug"],
		CategoryStore: h.CategoryStore,
	})
	if err != nil {
		log.Println("fetch category by slug:", err)
		HandleActionError(w, r, err)
		return
	}
	cat := res.(core.FetchCategoryBySlugRes).Category

	res, err = core.Do(core.FetchCategoryArticleList{
		CategoryID:    cat.ID.Hex(),
		CategoryStore: h.CategoryStore,
		ArticleStore:  h.ArticleStore,
	})
	if err != nil {
		log.Println("fetch article list:", err)
		HandleActionError(w, r, err)
		return
	}
	arts := res.(core.FetchCategoryArticleListRes).Articles

	err = ExecuteTemplate(CategoryViewTpl, w, struct {
		Page
		Context  Context
		Category *data.Category
		Articles []data.Article
	}{
		Page: Page{
			Title: cat.Title + " | Toph Help",
		},
		Context:  ctx,
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

	err = ExecuteTemplate(BackCategoryListTpl, w, struct {
		Page
		Context    Context
		Categories []data.Category
	}{
		Page: Page{
			Title: "Categories" + " | Toph Help",
		},
		Context:    ctx,
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
	ctx := GetContext(r)

	err := ExecuteTemplate(BackCategoryNewFormTpl, w, struct {
		Page
		Context Context
	}{
		Page: Page{
			Title: "Create New Category" + " | Toph Help",
		},
		Context: ctx,
	})
	if err != nil {
		log.Println("execute template:", err)
		HandleActionError(w, r, err)
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
		HandleActionError(w, r, err)
		return
	}

	_, err = core.Do(core.CreateCategory{
		Title:         body.Title,
		Slug:          body.Slug,
		Order:         body.Order,
		CategoryStore: h.CategoryStore,
	})
	if err != nil {
		log.Println("create category:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/_/categories", http.StatusSeeOther)
}

var BackCategoryEditFormTpl = template.Must(template.ParseFiles("ui/gohtml/layout.gohtml", "ui/gohtml/backcategoryeditform.gohtml", "ui/gohtml/backtabset.gohtml"))

type ServeBackCategoryEditForm struct {
	CategoryStore db.CategoryStore
}

func (h ServeBackCategoryEditForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := GetContext(r)

	vars := mux.Vars(r)

	res, err := core.Do(core.FetchCategory{
		ID:            vars["category_id"],
		CategoryStore: h.CategoryStore,
	})
	if err != nil {
		log.Println("fetch category:", err)
		HandleActionError(w, r, err)
		return
	}
	cat := res.(core.FetchCategoryRes).Category

	err = ExecuteTemplate(BackCategoryEditFormTpl, w, struct {
		Page
		Context  Context
		Category *data.Category
	}{
		Page: Page{
			Title: cat.Title + " | Toph Help",
		},
		Context:  ctx,
		Category: cat,
	})
	if err != nil {
		log.Println("execute template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

type HandleBackCategoryEditForm struct {
	CategoryStore db.CategoryStore
}

type HandleBackCategoryEditFormVal struct {
	Title string `schema:"title"`
	Slug  string `schema:"slug"`
	Order int    `schema:"order"`
}

func (h HandleBackCategoryEditForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	body := HandleBackCategoryEditFormVal{}
	err := ParseRequestBody(r, &body)
	if err != nil {
		log.Println("parse request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	_, err = core.Do(core.UpdateCategory{
		ID:            vars["category_id"],
		Title:         body.Title,
		Slug:          body.Slug,
		Order:         body.Order,
		CategoryStore: h.CategoryStore,
	})
	if err != nil {
		log.Println("update category:", err)
		HandleActionError(w, r, err)
		return
	}

	http.Redirect(w, r, "/_/categories", http.StatusSeeOther)
}
