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

	res, err = core.Do(core.FetchArticleBySlug{
		CategoryID:    cat.ID.Hex(),
		Slug:          vars["article_slug"],
		ArticleStore:  h.ArticleStore,
		CategoryStore: h.CategoryStore,
	})
	if err != nil {
		log.Println("fetch article by slug:", err)
		HandleActionError(w, r, err)
		return
	}
	art := res.(core.FetchArticleBySlugRes).Article

	err = ExecuteTemplate(ArticleViewTpl, w, struct {
		Context  Context
		Article  *data.Article
		Category *data.Category
	}{
		Context:  ctx,
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
	ctx := GetContext(r)

	res, err := core.Do(core.FetchArticleList{
		ArticleStore: h.ArticleStore,
	})
	if err != nil {
		log.Println("fetch article list:", err)
		HandleActionError(w, r, err)
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
			HandleActionError(w, r, err)
			return
		}
		artCat[art.ID.Hex()] = res.(core.FetchCategoryRes).Category
	}

	err = ExecuteTemplate(BackArticleListTpl, w, struct {
		Context         Context
		Articles        []data.Article
		ArticleCategory map[string]*data.Category
	}{
		Context:         ctx,
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

	err = ExecuteTemplate(BackArticleNewFormTpl, w, struct {
		Context    Context
		Categories []data.Category
	}{
		Context:    ctx,
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
		HandleActionError(w, r, err)
		return
	}

	http.Redirect(w, r, "/_/articles", http.StatusSeeOther)
}

var BackArticleEditFormTpl = template.Must(template.ParseFiles("ui/gohtml/layout.gohtml", "ui/gohtml/backarticleeditform.gohtml", "ui/gohtml/backtabset.gohtml"))

type ServeBackArticleEditForm struct {
	ArticleStore  db.ArticleStore
	CategoryStore db.CategoryStore
}

func (h ServeBackArticleEditForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := GetContext(r)

	vars := mux.Vars(r)

	res, err := core.Do(core.FetchCategoryList{
		CategoryStore: h.CategoryStore,
	})
	if err != nil {
		log.Println("fetch category list:", err)
		HandleActionError(w, r, err)
		return
	}
	cats := res.(core.FetchCategoryListRes).Categories

	res, err = core.Do(core.FetchArticle{
		ID:           vars["article_id"],
		ArticleStore: h.ArticleStore,
	})
	if err != nil {
		log.Println("fetch article:", err)
		HandleActionError(w, r, err)
		return
	}
	art := res.(core.FetchArticleRes).Article

	err = ExecuteTemplate(BackArticleEditFormTpl, w, struct {
		Context    Context
		Article    *data.Article
		Categories []data.Category
	}{
		Context:    ctx,
		Categories: cats,
		Article:    art,
	})
	if err != nil {
		log.Println("execute template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

type HandleBackArticleEditForm struct {
	ArticleStore  db.ArticleStore
	CategoryStore db.CategoryStore
}

type HandleBackArticleEditFormVal struct {
	CategoryID string `schema:"category_id"`
	Title      string `schema:"title"`
	Slug       string `schema:"slug"`
	Order      int    `schema:"order"`
	Content    string `schema:"content"`
}

func (h HandleBackArticleEditForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	body := HandleBackArticleEditFormVal{}
	err := ParseRequestBody(r, &body)
	if err != nil {
		log.Println("parse request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	_, err = core.Do(core.UpdateArticle{
		ID:            vars["article_id"],
		CategoryID:    body.CategoryID,
		Title:         body.Title,
		Slug:          body.Slug,
		Order:         body.Order,
		Content:       body.Content,
		ArticleStore:  h.ArticleStore,
		CategoryStore: h.CategoryStore,
	})
	if err != nil {
		log.Println("update article:", err)
		HandleActionError(w, r, err)
		return
	}

	http.Redirect(w, r, "/_/articles", http.StatusSeeOther)
}
