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

var BackArticleNewFormTpl = template.Must(template.ParseFiles("ui/gohtml/layout.gohtml", "ui/gohtml/backarticlenewform.gohtml", "ui/gohtml/backtabset.gohtml"))

type ServeBackArticleNewForm struct {
	CategoryStore db.CategoryStore
}

func (h ServeBackArticleNewForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := core.Do(core.FetchCategoryList{
		CategoryStore: h.CategoryStore,
	})
	if err != nil {
		log.Println("fetch category list:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	cats := res.(core.FetchCategoryListRes).Categories

	err = BackArticleNewFormTpl.Execute(w, struct {
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

type HandleBackArticleNewForm struct {
	ArticleStore  db.ArticleStore
	CategoryStore db.CategoryStore
}

type HandleBackArticleNewFormVal struct {
	CategoryID string `schema:"category_id"`
	Title      string `schema:"title"`
	Slug       string `schema:"slug"`
	Order      int    `schema:"order"`
	Content    string `schema:"content"`
}

func (h HandleBackArticleNewForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := HandleBackArticleNewFormVal{}
	err := ParseRequestBody(r, &body)
	if err != nil {
		log.Println("parse request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	_, err = core.Do(core.CreateArticle{
		CategoryID:    body.CategoryID,
		Title:         body.Title,
		Slug:          body.Slug,
		Order:         body.Order,
		Content:       body.Content,
		ArticleStore:  h.ArticleStore,
		CategoryStore: h.CategoryStore,
	})
	if err != nil {
		log.Println("create article:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/_/articles", http.StatusSeeOther)
}
