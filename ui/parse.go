package ui

import (
	"net/http"

	"github.com/gorilla/schema"
)

func ParseRequestBody(r *http.Request, v interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	return schema.NewDecoder().Decode(v, r.Form)
}
