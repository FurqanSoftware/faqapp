package ui

import (
	"net/http"

	"git.furqan.io/faqapp/faqapp/data"

	"github.com/gorilla/context"
)

type Context struct {
	Account  *data.Account
	Settings map[string]interface{}
}

func GetContext(r *http.Request) Context {
	return context.Get(r, "context").(Context)
}
