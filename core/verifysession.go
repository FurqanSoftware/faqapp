package core

import (
	"time"

	"git.furqan.io/faqapp/faqapp/db"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

type VerifySession struct {
	Claims jwt.MapClaims

	AccountRepo db.Accounts
}

func (a VerifySession) Do() (res Result, err error) {
	if a.Claims == nil {
		return VerifySessionRes{
			Okay: false,
		}, nil
	}
	expireAt, ok := a.Claims["expire_at"].(float64)
	if time.Unix(int64(expireAt), 0).Before(time.Now()) {
		return VerifySessionRes{
			Okay: false,
		}, nil
	}
	idStr, ok := a.Claims["id"].(string)
	if !ok || !bson.IsObjectIdHex(idStr) {
		return VerifySessionRes{
			Okay: false,
		}, nil
	}
	acc, err := a.AccountRepo.Get(bson.ObjectIdHex(idStr))
	if err != nil {
		return nil, DatabaseError{"VerifySession", err}
	}
	if acc == nil {
		return VerifySessionRes{
			Okay: false,
		}, nil
	}
	return VerifySessionRes{
		Okay: true,
	}, nil
}

type VerifySessionRes struct {
	Okay bool
}
