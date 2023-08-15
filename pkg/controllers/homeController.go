package controllers

import (
	"log"
	"mvc/pkg/models"
	"mvc/pkg/views"

	"net/http"
)

func ViewHome(w http.ResponseWriter, r *http.Request) {

	userId := r.Context().Value("userId").(int)
	admin := r.Context().Value("admin").(bool)

	if admin {
		data, err := models.GetDataAdmin()
		views.RenderTemplate(w, "home-admin", data)
		if err != nil {
			log.Fatal(err)
		}

	} else {

		data, err := models.GetDataUser(userId, "")
		views.RenderTemplate(w, "home", data)
		if err != nil {
			log.Fatal(err)
		}
	}
}