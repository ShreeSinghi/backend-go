package controllers

import (
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
