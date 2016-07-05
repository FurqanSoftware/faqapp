package data

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Category struct {
	ID bson.ObjectId

	Title string
	Order int

	CreatedAt  time.Time
	ModifiedAt time.Time
}
