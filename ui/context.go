package ui

import (
	"net/http"

	"git.furqansoftware.net/faqapp/faqapp/data"
)

type Context struct {
	Account  *data.Account
	Settings map[string]interface{}
}

func GetContext(r *http.Request) Context {
	return r.Context().Value("context").(Context)
}
