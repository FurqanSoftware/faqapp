package ui

import (
	"html/template"
	"log"
	"net/http"

	"git.furqansoftware.net/faqapp/faqapp/core"
	"git.furqansoftware.net/faqapp/faqapp/data"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

var BackSettingPasswordFormTpl = template.Must(template.ParseFiles("ui/gohtml/layout.gohtml", "ui/gohtml/backsettingpasswordform.gohtml", "ui/gohtml/backtabset.gohtml", "ui/gohtml/backsettingsnavset.gohtml"))

type ServeBackSettingPasswordForm struct{}

func (h ServeBackSettingPasswordForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := GetContext(r)

	err := ExecuteTemplate(BackSettingPasswordFormTpl, w, struct {
		Context Context
	}{
		Context: ctx,
	})
	if err != nil {
		log.Println("execute template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

type HandleBackSettingPasswordForm struct {
	AccountStore db.AccountStore
}

type HandleBackSettingPasswordFormVal struct {
	Current string `schema:"current"`
	New     string `schema:"new"`
	Confirm string `schema:"confirm"`
}

func (h HandleBackSettingPasswordForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := GetContext(r)

	body := HandleBackSettingPasswordFormVal{}
	err := ParseRequestBody(r, &body)
	if err != nil {
		log.Println("parse request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	_, err = core.Do(core.UpdateAccountPassword{
		ID:           ctx.Account.ID.Hex(),
		Current:      body.Current,
		New:          body.New,
		Confirm:      body.Confirm,
		AccountStore: h.AccountStore,
	})
	if err != nil {
		log.Println("update account password:", err)
		HandleActionError(w, r, err)
		return
	}

	http.Redirect(w, r, "/_/settings/password", http.StatusSeeOther)
}

var BackSettingAdvancedFormTpl = template.Must(template.ParseFiles("ui/gohtml/layout.gohtml", "ui/gohtml/backsettingadvancedform.gohtml", "ui/gohtml/backtabset.gohtml", "ui/gohtml/backsettingsnavset.gohtml"))

type ServeBackSettingAdvancedForm struct {
	SettingStore db.SettingStore
}

func (h ServeBackSettingAdvancedForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := GetContext(r)

	res, err := core.Do(core.FetchSettingList{
		SettingStore: h.SettingStore,
	})
	if err != nil {
		log.Println("fetch setting list:", err)
		HandleActionError(w, r, err)
		return
	}
	stts := res.(core.FetchSettingListRes).Settings

	err = ExecuteTemplate(BackSettingAdvancedFormTpl, w, struct {
		Context  Context
		Settings []data.Setting
	}{
		Context:  ctx,
		Settings: stts,
	})
	if err != nil {
		log.Println("execute template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

type HandleBackSettingAdvancedForm struct {
	SettingStore db.SettingStore
}

type HandleBackSettingAdvancedFormVal struct {
	Settings []struct {
		Key   string `schema:"key"`
		Value string `schema:"value"`
	} `schema:"settings"`
}

func (h HandleBackSettingAdvancedForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := HandleBackSettingAdvancedFormVal{}
	err := ParseRequestBody(r, &body)
	if err != nil {
		log.Println("parse request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	for _, stt := range body.Settings {
		if stt.Key == "" {
			continue
		}

		_, err = core.Do(core.UpdateSetting{
			Key:          stt.Key,
			Value:        stt.Value,
			SettingStore: h.SettingStore,
		})
		if err != nil {
			log.Println("update setting bulk:", err)
			HandleActionError(w, r, err)
			return
		}
	}

	http.Redirect(w, r, "/_/settings/advanced", http.StatusSeeOther)
}
