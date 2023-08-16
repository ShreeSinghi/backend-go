package controllers

import (
	"fmt"
	"log"
	"mvc/pkg/models"
	"mvc/pkg/views"
	"strings"

	"net/http"
)

func ViewLogin(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "login", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("rghrthr")
	r.ParseForm()
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	data := struct {
		Error string
	}{
		Error: "Invalid username or password",
	}

	fmt.Println(username, password)

	sessionString, err, success := models.Login(username, password)
	if err != nil {
		log.Fatal(err)
	}

	if !success {
		views.RenderTemplate(w, "login", data)
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     "sessionID",
			Value:    sessionString,
			HttpOnly: true,
		})

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}

}

func Logout(w http.ResponseWriter, r *http.Request) {
	authorised := r.Context().Value("authorised").(bool)
	log.Println("hey")

	log.Println("hasdasey")

	oldcookie := r.Header.Get("Cookie")
	cookieid := oldcookie[strings.Index(oldcookie, "sessionID=")+10:]
	models.Logout(cookieid)

	cookie := http.Cookie{
		Name:     "sessionID",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)

}
