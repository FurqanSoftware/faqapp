package ui

import (
	"io"
	"log"
	"net/http"

	"git.furqansoftware.net/faqapp/faqapp/core"
	"git.furqansoftware.net/faqapp/faqapp/db"
)

type ServeCustomCSS struct {
	SettingStore db.SettingStore
}

func (h ServeCustomCSS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := core.Do(core.FetchSetting{
		Key:          "appearance.custom_css",
		SettingStore: h.SettingStore,
	})
	if err != nil {
		log.Println("create session:", err)
		HandleActionError(w, r, err)
		return
	}
	stt := res.(core.FetchSettingRes).Setting

	w.Header().Set("Content-Type", "text/css")
	if stt == nil {
		return
	}
	io.WriteString(w, stt.ValueString())
}
