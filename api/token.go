package api

import (
	"fmt"
	"net/http"
	"strings"

	"git.furqan.io/faqapp/faqapp/cfg"

	"github.com/dgrijalva/jwt-go"
)

func GetRequestClaims(r *http.Request) (jwt.MapClaims, error) {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Token" {
		return nil, nil
	}

	token, err := jwt.ParseWithClaims(auth[1], jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
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
