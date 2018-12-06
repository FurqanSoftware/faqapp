package api

import (
	"log"
	"net/http"
	"time"

	"git.furqansoftware.net/faqapp/faqapp/core"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

type ServeCategoryList struct {
	CategoryStore db.CategoryStore
}

type ServeCategoryListRes []ServeCategoryListResItem

type ServeCategoryListResItem struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Order      int       `json:"order"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func (h ServeCategoryList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := core.Do(core.FetchCategoryList{
		CategoryStore: h.CategoryStore,
	})
	if err != nil {
		log.Println("fetch category list:", err)
		HandleActionError(w, r, err)
		return
	}

	resp := ServeCategoryListRes{}
	for _, cat := range res.(core.FetchCategoryListRes).Categories {
		resp = append(resp, ServeCategoryListResItem{
			ID:         cat.ID.Hex(),
			Title:      cat.Title,
			Slug:       cat.Slug,
			Order:      cat.Order,
			CreatedAt:  cat.CreatedAt,
			ModifiedAt: cat.ModifiedAt,
		})
	}
	ServeResult(w, r, resp)
}

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
