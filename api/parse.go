package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseRequestBody(r *http.Request, v interface{}) error {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		return parseJSONBody(r.Body, v)

	default:
		return parseJSONBody(r.Body, v)
	}
}

func parseJSONBody(body io.Reader, v interface{}) error {
	return json.NewDecoder(body).Decode(v)
}
