package controllers

import (
	"log"
	"mvc/pkg/models"
	"net/http"
	"strconv"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	// Fetch user authentication data from the context
	admin := r.Context().Value("admin").(bool)
	title := r.FormValue("title")
	quantityStr := r.FormValue("quantity")
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}

	if quantity < 0 || title == "" {
		http.Error(w, "No negative values or empty string allowed", http.StatusBadRequest)
		return
	}

	if !admin {
		http.Error(w, "Not authenticated", http.StatusForbidden)
		return
	}

	err = models.AddBook(title, quantity)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func ProcessChecks(w http.ResponseWriter, r *http.Request) {
	// Fetch user authentication data from the context
	admin := r.Context().Value("admin").(bool)

	db, err := models.Connection()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	r.ParseForm()

	checkRequests := r.PostForm
	if !admin {
		http.Error(w, "Not authenticated", http.StatusForbidden)
		return
	}

	err = models.ProcessChecks(checkRequests)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func ProcessAdminRequests(w http.ResponseWriter, r *http.Request) {
	// Fetch user authentication data from the context
	admin := r.Context().Value("admin").(bool)

	if !admin {
		http.Error(w, "Not authenticated", http.StatusForbidden)
		return
	}

	r.ParseForm()

	requestedUsers := r.PostForm
	log.Println(requestedUsers)
	err := models.ProcessAdminRequests(requestedUsers)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/home-admin", http.StatusSeeOther)
}
