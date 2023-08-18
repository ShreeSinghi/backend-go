package controllers

import (
	"log"
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
)

func ViewRegister(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "register", nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	data := struct {
		Error string
	}{
		Error: "User already exists",
	}

	if username == "" || password == "" {
		http.Error(w, "Empty username or password", http.StatusBadRequest)
		return
	}

	hash, err := models.HashKaro(password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	_, err = models.RegisterUser(username, hash)
	if err != nil {
		views.RenderTemplate(w, "register", data)
		return
	}

	sessionString, err, _ := models.Login(username, password)
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
