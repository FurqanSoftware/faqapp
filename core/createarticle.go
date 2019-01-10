package core

import (
	"git.furqansoftware.net/faqapp/faqapp/data"
	"git.furqansoftware.net/faqapp/faqapp/db"
	"gopkg.in/mgo.v2/bson"
)

type CreateArticle struct {
	CategoryID string
	Title      string
	Slug       string
	Order      int
	Content    string
	Published  bool

	ArticleStore  db.ArticleStore
	CategoryStore db.CategoryStore
}

func (a CreateArticle) Validate() error {
	if !bson.IsObjectIdHex(a.CategoryID) {
		return ValidationError{"CreateArticle", "CategoryID", IssueInvalid}
	}
	if a.Title == "" {
		return ValidationError{"CreateArticle", "Title", IssueMissing}
	}
	if len(a.Title) > 256 {
		return ValidationError{"CreateArticle", "Title", IssueTooLong}
	}
	if a.Slug == "" {
		return ValidationError{"CreateArticle", "Slug", IssueMissing}
	}
	if len(a.Slug) > 256 {
		return ValidationError{"CreateArticle", "Slug", IssueTooLong}
	}
	return nil
}

func (a CreateArticle) Do() (res Result, err error) {
	cat, err := a.CategoryStore.Get(bson.ObjectIdHex(a.CategoryID))
	if err != nil {
		return nil, DatabaseError{"CreateArticle", err}
	}

	art := data.Article{
		CategoryID: cat.ID,
		Title:      a.Title,
		Slug:       a.Slug,
		Order:      a.Order,
		Published:  a.Published,
	}
	art.SetContent(a.Content)
	art.SetPublishedAt(a.Published)
	err = a.ArticleStore.Put(&art)
	if err != nil {
		return nil, DatabaseError{"CreateArticle", err}
	}
	return CreateArticleRes{
		Article: &art,
	}, nil
}

type CreateArticleRes struct {
	Article *data.Article
}
