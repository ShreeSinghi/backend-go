package controllers

import (
	"net/http"
	"fmt"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	authorised := r.Context().Value("authorised").(bool)
	fmt.Println(authorised)
	if !authorised {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}

}
