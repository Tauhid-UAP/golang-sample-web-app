package handlers

import (
	"html/template"
	"net/http"


	"github.com/Tauhid-UAP/golang-sample-web-app/core/models"
)

var templates = template.Must(template.ParseGlob("core/templates/*.html"),)

type PageData struct {
	Title string
	User models.User
	CSRF string
	StaticAssetBaseURL string
}

func Render(w http.ResponseWriter, page string, data PageData) {
	template, err := template.ParseFiles(
		"core/templates/base.html",
		"core/templates/"+page,
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = template.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
