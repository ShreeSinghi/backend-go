package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"mvc/pkg/models"
	"net/http"
	"strconv"
)

func RequestCheckout(w http.ResponseWriter, r *http.Request) {

	var requestBody struct {
		bookId string `json:"bookId, string"`
	}

	userId := r.Context().Value("userId").(int)

	r.ParseForm()
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Fatal(requestBody)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	bookId := requestBody.bookId

	fmt.Println(bookId)
	bookIdint, err := strconv.Atoi(bookId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := models.RequestCheckout(bookIdint, userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func ReturnBook(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(int)

	bookIdStr := r.FormValue("bookId")
	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	err = models.ReturnBook(bookId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Book return requested")
}

func RequestAdmin(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(int)

	fmt.Println("yoo")

	err := models.RequestAdmin(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := models.GetDataUser(userId, "")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}
