package ui

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func GetSession(store sessions.Store, r *http.Request, name string) (*sessions.Session, error) {
	sess, err := store.Get(r, name)
	if err != nil {
		sess, _ = store.New(r, name)
	}
	return sess, nil
}
