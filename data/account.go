package data

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID bson.ObjectId

	Handle   string
	Password AccountPassword

	FirstIP  string
	RecentIP string

	CreatedAt  time.Time
	ModifiedAt time.Time
}
