package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
	"strconv"
)

func ViewRequestReturn(w http.ResponseWriter, r *http.Request) {

	userId := r.Context().Value("userId").(int)

	data, err := models.GetDataUser(userId, "")
	views.RenderTemplate(w, "request-return", data)
	if err != nil {
		log.Fatal(err)
	}

}

func RequestCheckout(w http.ResponseWriter, r *http.Request) {

	var requestBody struct {
		BookId string `json:"bookId"`
	}

	userId := r.Context().Value("userId").(int)

	r.ParseForm()
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Fatal(requestBody)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	bookId := requestBody.BookId

	fmt.Println(bookId, "meow")
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

func RequestCheckin(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(int)
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bookId, err := strconv.Atoi(requestBody["bookId"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	err = models.RequestCheckin(bookId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Book return requested")
}

func RequestAdmin(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(int)

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
