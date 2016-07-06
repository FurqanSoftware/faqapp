package data

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID bson.ObjectId `bson:"_id,omitempty"`

	Handle   string          `bson:"handle"`
	Password AccountPassword `bson:"password"`

	FirstIP  string `bson:"first_ip"`
	RecentIP string `bson:"recent_ip"`

	CreatedAt  time.Time `bson:"created_at"`
	ModifiedAt time.Time `bson:"modified_at"`
}

func (a *Account) PreCreate() {
	a.ID = bson.NewObjectId()
	a.CreatedAt = time.Now()
	a.ModifiedAt = a.CreatedAt
}

func (a *Account) PreModify() {
	a.ModifiedAt = time.Now()
}
