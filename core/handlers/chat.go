package handlers

import (
	"net/http"
	"github.com/Tauhid-UAP/golang-sample-web-app/core/middleware"
)

func Chat(w http.ResponseWriter, r *http.Request) {
	Render(w, "chat.html", PageData{
		Title: "Global Chat",
		CSRF: r.Context().Value(middleware.CSRFKey).(string),
	})
}
