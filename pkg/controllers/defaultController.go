package controllers

import (
	"mvc/pkg/views"
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "404", nil)
}
