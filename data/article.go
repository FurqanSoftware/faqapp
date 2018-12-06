package data

import (
	"html/template"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"

	"gopkg.in/mgo.v2/bson"
)

type Article struct {
	ID bson.ObjectId `bson:"id"`

	CategoryID bson.ObjectId `bson:"category_id"`

	Title string `bson:"title"`
	Slug  string `bson:"slug"`

	Order int `bson:"order"`

	Content     string        `bson:"content"`
	ContentHTML template.HTML `bson:"content_html"`

	CreatedAt  time.Time `bson:"created_at"`
	ModifiedAt time.Time `bson:"modified_at"`
}

func (a *Article) SetContent(content string) {
	a.Content = content
	a.ContentHTML = template.HTML(bluemonday.UGCPolicy().SanitizeBytes(blackfriday.Run([]byte(content))))
}

func (a *Article) PreCreate() {
	a.ID = bson.NewObjectId()
	a.CreatedAt = time.Now()
	a.ModifiedAt = a.CreatedAt
}

func (a *Article) PreModify() {
	a.ModifiedAt = time.Now()
}
