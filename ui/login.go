package ui

import (
	"html/template"
	"log"
	"net/http"

	"git.furqan.io/faqapp/faqapp/core"
	"git.furqan.io/faqapp/faqapp/db"

	"github.com/gorilla/sessions"
)

var LoginFormTpl = template.Must(template.ParseFiles("ui/gohtml/layout.gohtml", "ui/gohtml/loginform.gohtml"))

type ServeLoginForm struct {
}

func (h ServeLoginForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := GetContext(r)

	err := ExecuteTemplate(LoginFormTpl, w, struct {
		Context Context
	}{
		Context: ctx,
	})
	if err != nil {
		log.Println("execute template:", err)
		HandleActionError(w, r, err)
		return
	}
}

type HandleLoginForm struct {
	AccountStore db.AccountStore
	SessionStore sessions.Store
}

type HandleLoginFormVal struct {
	Handle   string `schema:"handle"`
	Password string `schema:"password"`
}

func (h HandleLoginForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := HandleLoginFormVal{}
	err := ParseRequestBody(r, &body)
	if err != nil {
		log.Println("parse request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	res, err := core.Do(core.CreateSession{
		Handle:       body.Handle,
		Password:     body.Password,
		AccountStore: h.AccountStore,
	})
	if err != nil {
		log.Println("create session:", err)
		HandleActionError(w, r, err)
		return
	}

	csRes := res.(core.CreateSessionRes)
	if !csRes.Match {
		http.Redirect(w, r, "/_", http.StatusSeeOther)
		return
	}

	sess, err := GetSession(h.SessionStore, r, "s")
	if err != nil {
		log.Println("get session:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	sess.Values["token"] = csRes.Token
	err = sess.Save(r, w)
	if err != nil {
		log.Println("save session:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/_", http.StatusSeeOther)
}

type HandleLogout struct {
	SessionStore sessions.Store
}

func (h HandleLogout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess, err := GetSession(h.SessionStore, r, "s")
	if err != nil {
		log.Println("get session:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	delete(sess.Values, "token")
	err = sess.Save(r, w)
	if err != nil {
		log.Println("save session:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
