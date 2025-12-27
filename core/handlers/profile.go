package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"log"

	"github.com/Tauhid-UAP/golang-sample-web-app/core/middleware"
	"github.com/Tauhid-UAP/golang-sample-web-app/core/store"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)
	user, _ := store.GetUserByID(r.Context(), userID)

	if r.Method == http.MethodPost {
		user.FirstName = r.FormValue("FirstName")
		user.LastName = r.FormValue("LastName")

		file, header, err := r.FormFile("ProfileImage")
		if err != nil {
			log.Printf("File upload error: %v", err)
			store.UpdateUser(r.Context(), user)
			http.Redirect(w, r, "/profile", http.StatusSeeOther)
			return
		}

		defer file.Close()
		path := filepath.Join("media/uploads", userID+"_"+header.Filename)
		destination, _ := os.Create(path)
		defer destination.Close()

		io.Copy(destination, file)
		user.ProfileImage = &path		

		store.UpdateUser(r.Context(), user)
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	}

	Render(w, "profile.html", PageData{
		Title: "Profile",
		User: user,
		CSRF: r.Context().Value(middleware.CSRFKey).(string),
	})
}
