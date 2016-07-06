package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"git.furqan.io/faqapp/faqapp/core"
)

func HandleActionError(w http.ResponseWriter, r *http.Request, err error) {
	switch err := err.(type) {
	case core.ValidationError:
		ServeError(w, r, err.Element+" is "+string(err.Issue), http.StatusBadRequest)

	default:
		ServeInternalServerError(w, r)
	}
}

func ServeUnauthorized(w http.ResponseWriter, r *http.Request) {
	ServeError(w, r, "Unauthorized", http.StatusUnauthorized)
}

func ServeBadRequest(w http.ResponseWriter, r *http.Request) {
	ServeError(w, r, "Bad Request", http.StatusBadRequest)
}

func ServeInternalServerError(w http.ResponseWriter, r *http.Request) {
	ServeError(w, r, "Internal Server Error", http.StatusInternalServerError)
}

func ServeError(w http.ResponseWriter, r *http.Request, error string, code int) {
	switch r.Header.Get("Accept") {
	case "application/json":
		serveJSONError(w, r, error, code)

	default:
		serveJSONError(w, r, error, code)
	}
}

func serveJSONError(w http.ResponseWriter, r *http.Request, error string, code int) {
	b := &bytes.Buffer{}
	err := json.NewEncoder(b).Encode(map[string]interface{}{
		"error": error,
	})
	catch(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = io.Copy(w, b)
	catch(err)
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}
