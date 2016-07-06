package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"git.furqan.io/faqapp/faqapp/core"
	"git.furqan.io/faqapp/faqapp/db"
)

type CreateArticle struct {
	ArticleStore  db.ArticleStore
	CategoryStore db.CategoryStore
}

type CreateArticleVal struct {
	CategoryID string `json:"category_id"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Order      int    `json:"order"`
	Content    string `json:"content"`
}

type CreateArticleRes struct {
	ID string `json:"id"`
}

func (h CreateArticle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	body := CreateArticleVal{}
	err := ParseRequestBody(r, &body)
	if err != nil {
		log.Println("parse request body:", err)
		ServeBadRequest(w, r)
		return
	}

	res, err := core.Do(core.CreateArticle{
		CategoryID:    vars["category_id"],
		Title:         body.Title,
		Slug:          body.Slug,
		Order:         body.Order,
		Content:       body.Content,
		ArticleStore:  h.ArticleStore,
		CategoryStore: h.CategoryStore,
	})
	if err != nil {
		log.Println("create category:", err)
		HandleActionError(w, r, err)
		return
	}

	resp := CreateArticleRes{
		ID: res.(core.CreateArticleRes).Article.ID.Hex(),
	}
	ServeResult(w, r, resp)
}
