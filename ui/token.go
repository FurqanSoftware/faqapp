package ui

import (
	"fmt"
	"net/http"

	"git.furqan.io/faqapp/faqapp/cfg"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
)

func GetRequestClaims(store sessions.Store, r *http.Request) (jwt.MapClaims, error) {
	sess, err := GetSession(store, r, "s")
	if err != nil {
		return nil, err
	}
	jwtStr, ok := sess.Values["token"].(string)
	if !ok {
		return nil, nil
	}

	token, err := jwt.ParseWithClaims(jwtStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil
	}
	return claims, nil
}
