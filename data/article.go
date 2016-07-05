package data

import (
	"html/template"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Article struct {
	ID bson.ObjectId

	CategoryID string

	Title string
	Order int

	Content     string
	ContentHTML template.HTML

	CreatedAt  time.Time
	ModifiedAt time.Time
}
