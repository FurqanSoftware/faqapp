package data

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Category struct {
	ID bson.ObjectId

	Title string
	Slug  string

	Order int

	CreatedAt  time.Time
	ModifiedAt time.Time
}

func (c *Category) PreCreate() {
	c.ID = bson.NewObjectId()
	c.CreatedAt = time.Now()
	c.ModifiedAt = c.CreatedAt
}

func (c *Category) PreModify() {
	c.ModifiedAt = time.Now()
}
