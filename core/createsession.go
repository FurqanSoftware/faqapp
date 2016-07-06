package core

import (
	"time"

	"git.furqan.io/faqapp/faqapp/cfg"
	"git.furqan.io/faqapp/faqapp/db"
	"github.com/dgrijalva/jwt-go"
)

type CreateSession struct {
	Handle   string
	Password string

	AccountRepo db.Accounts
}

func (a CreateSession) Do() (res Result, err error) {
	acc, err := a.AccountRepo.GetByHandle(a.Handle)
	if err != nil {
		return nil, DatabaseError{"CreateSession", err}
	}
	if !acc.Password.Match(a.Password) {
		return CreateSessionRes{
			Match: false,
			Token: "",
		}, nil
	}

	token := jwt.New(jwt.SigningMethodHS256)
	now := time.Now()
	token.Claims = jwt.MapClaims{
		"id":         acc.ID.Hex(),
		"created_at": now.Unix(),
		"expire_at":  now.Add(24 * time.Hour).Unix(),
	}
	jwtStr, err := token.SignedString([]byte(cfg.Secret))
	if err != nil {
		return nil, err
	}

	return CreateSessionRes{
		Match: true,
		Token: jwtStr,
	}, nil
}

type CreateSessionRes struct {
	Match bool
	Token string
}
