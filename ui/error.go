package ui

import (
	"net/http"

	"git.furqan.io/faqapp/faqapp/core"
	"git.furqan.io/faqapp/faqapp/db"
)

func HandleActionError(w http.ResponseWriter, r *http.Request, err error) {
	switch err := err.(type) {
	case core.ValidationError:
		http.Error(w, err.Element+" is "+string(err.Issue), http.StatusBadRequest)

	case core.DatabaseError:
		if err.Base == db.ErrNotFound {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		ServeInternalServerError(w, r)

	default:
		ServeInternalServerError(w, r)
	}
}

func ServeUnauthorized(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func ServeBadRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Bad Request", http.StatusBadRequest)
}

func ServeInternalServerError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
