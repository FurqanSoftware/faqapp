package ui

import (
	"html/template"
	"log"
	"net/http"

	"git.furqan.io/faqapp/faqapp/core"
	"git.furqan.io/faqapp/faqapp/data"
	"git.furqan.io/faqapp/faqapp/db"
)

var BackSettingFormTpl = template.Must(template.ParseFiles("ui/gohtml/layout.gohtml", "ui/gohtml/backsettingform.gohtml", "ui/gohtml/backtabset.gohtml"))

type ServeBackSettingForm struct {
	SettingStore db.SettingStore
}

func (h ServeBackSettingForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := core.Do(core.FetchSettingList{
		SettingStore: h.SettingStore,
	})
	if err != nil {
		log.Println("fetch setting list:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	stts := res.(core.FetchSettingListRes).Settings

	err = ExecuteTemplate(BackSettingFormTpl, w, struct {
		Settings []data.Setting
	}{
		Settings: stts,
	})
	if err != nil {
		log.Println("execute template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

type HandleBackSettingForm struct {
	SettingStore db.SettingStore
}

type HandleBackSettingFormVal struct {
	Settings []struct {
		Key   string `schema:"key"`
		Value string `schema:"value"`
	} `schema:"settings"`
}

func (h HandleBackSettingForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := HandleBackSettingFormVal{}
	err := ParseRequestBody(r, &body)
	if err != nil {
		log.Println("parse request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	for _, stt := range body.Settings {
		_, err = core.Do(core.UpdateSetting{
			Key:          stt.Key,
			Value:        stt.Value,
			SettingStore: h.SettingStore,
		})
		if err != nil {
			log.Println("update setting bulk:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/_/settings", http.StatusSeeOther)
}
