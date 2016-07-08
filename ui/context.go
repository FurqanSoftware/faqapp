package ui

import (
	"net/http"

	"github.com/gorilla/context"
)

type Context struct {
	Settings map[string]interface{}
}

func GetContext(r *http.Request) Context {
	return context.Get(r, "context").(Context)
}
