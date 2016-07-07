package ui

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
)

func ExecuteTemplate(tpl *template.Template, w http.ResponseWriter, v interface{}) error {
	b := bytes.Buffer{}
	err := tpl.Execute(&b, v)
	if err != nil {
		return err
	}
	_, err = io.Copy(w, &b)
	return err
}
