package core

import (
	"gopkg.in/mgo.v2/bson"

	"git.furqansoftware.net/faqapp/faqapp/data"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

type UpdateArticle struct {
	ID         string
	CategoryID string
	Title      string
	Slug       string
	Order      int
	Content    string
	Published  bool

	ArticleStore  db.ArticleStore
	CategoryStore db.CategoryStore
}

func (a UpdateArticle) Validate() error {
	if !bson.IsObjectIdHex(a.ID) {
		return ValidationError{"UpdateArticle", "ID", IssueInvalid}
	}
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

func (a UpdateArticle) Do() (Result, error) {
	cat, err := a.CategoryStore.Get(bson.ObjectIdHex(a.CategoryID))
	if err != nil {
		return nil, DatabaseError{"UpdateArticle", err}
	}

	art, err := a.ArticleStore.Get(bson.ObjectIdHex(a.ID))
	if err != nil {
		return nil, DatabaseError{"UpdateArticle", err}
	}

	art.CategoryID = cat.ID
	art.Title = a.Title
	art.Slug = a.Slug
	art.Order = a.Order
	art.SetContent(a.Content)
	art.SetPublished(a.Published)
	err = a.ArticleStore.Put(art)
	if err != nil {
		return nil, DatabaseError{"UpdateArticle", err}
	}

	return UpdateArticleRes{
		Article: art,
	}, nil
}

type UpdateArticleRes struct {
	Article *data.Article
}
