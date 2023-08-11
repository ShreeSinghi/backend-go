package controllers

import (
	"fmt"
	"log"
	"mvc/pkg/models"
	"mvc/pkg/views"

	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "login", nil)
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("rghrthr")
	r.ParseForm()
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	fmt.Println(username, password)

	sessionString, err := models.Login(username, password)
	if err != nil {
		log.Fatal(err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "sessionID",
		Value:    sessionString,
		HttpOnly: true,
	})

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
