package api

import (
	"log"
	"net/http"

	"git.furqan.io/faqapp/faqapp/core"
	"git.furqan.io/faqapp/faqapp/db"
)

type CreateCategory struct {
	CategoryStore db.CategoryStore
}

type CreateCategoryVal struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
	Order int    `json:"order"`
}

type CreateCategoryRes struct {
	ID string `json:"id"`
}

func (h CreateCategory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := CreateCategoryVal{}
	err := ParseRequestBody(r, &body)
	if err != nil {
		log.Println("parse request body:", err)
		ServeBadRequest(w, r)
		return
	}

	res, err := core.Do(core.CreateCategory{
		Title:         body.Title,
		Slug:          body.Slug,
		Order:         body.Order,
		CategoryStore: h.CategoryStore,
	})
	if err != nil {
		log.Println("create category:", err)
		HandleActionError(w, r, err)
		return
	}

	resp := CreateCategoryRes{
		ID: res.(core.CreateCategoryRes).Category.ID.Hex(),
	}
	ServeResult(w, r, resp)
}
