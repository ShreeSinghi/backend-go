package controllers

import (
	"log"
	"mvc/pkg/models"
	"mvc/pkg/views"

	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "login", nil)
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	newSessionID, err := models.Login(username, password)
	if err != nil {
		log.Fatal(err)
	}

	// http.SetCookie(w, &http.Cookie{
	// 	Name:     "sessionID",
	// 	Value:    fmt.Sprintf("%x", newSessionID),
	// 	HttpOnly: true,
	// })

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
