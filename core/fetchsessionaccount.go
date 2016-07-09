package core

import (
	"time"

	"git.furqan.io/faqapp/faqapp/data"
	"git.furqan.io/faqapp/faqapp/db"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

type FetchSessionAccount struct {
	Claims jwt.MapClaims

	AccountStore db.AccountStore
}

func (a FetchSessionAccount) Do() (res Result, err error) {
	if a.Claims == nil {
		return FetchSessionAccountRes{}, nil
	}
	expireAt, ok := a.Claims["expire_at"].(float64)
	if time.Unix(int64(expireAt), 0).Before(time.Now()) {
		return FetchSessionAccountRes{}, nil
	}
	idStr, ok := a.Claims["id"].(string)
	if !ok || !bson.IsObjectIdHex(idStr) {
		return FetchSessionAccountRes{}, nil
	}
	acc, err := a.AccountStore.Get(bson.ObjectIdHex(idStr))
	if err != nil {
		return nil, DatabaseError{"FetchSessionAccount", err}
	}
	if acc == nil {
		return FetchSessionAccountRes{}, nil
	}
	return FetchSessionAccountRes{
		Account: acc,
	}, nil
}

type FetchSessionAccountRes struct {
	Account *data.Account
}
