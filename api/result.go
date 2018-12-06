package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"git.furqansoftware.net/faqapp/faqapp/core"
)

func ServeResult(w http.ResponseWriter, r *http.Request, res core.Result) {
	switch r.Header.Get("Accept") {
	case "application/json":
		serveJSON(w, r, res)

	default:
		serveJSON(w, r, res)
	}
}

func serveJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	b := &bytes.Buffer{}
	err := json.NewEncoder(b).Encode(v)
	catch(err)

	w.Header().Set("Content-Type", "application/json")
	_, err = io.Copy(w, b)
	catch(err)
}
