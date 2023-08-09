package controllers

import (
	"mvc/pkg/models"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("views/*"))

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "register.gohtml", nil)
}

func RegisterPostHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Empty username or password", http.StatusBadRequest)
		return
	}

	hash, err := models.HashKaro(password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	userID, err := models.RegisterUser(username, hash)
	if err != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	err = models.CreateCookie(userID)
	if err != nil {
		http.Error(w, "Error creating cookie", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
