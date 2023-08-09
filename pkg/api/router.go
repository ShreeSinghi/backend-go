package api

import (
	"mvc/pkg/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()

	r.HandleFunc("/login", controllers.LoginHandler).Methods("GET")
	r.HandleFunc("/login", controllers.LoginPostHandler).Methods("POST")
	r.HandleFunc("/register", RegisterHandler).Methods("GET")
	r.HandleFunc("/register", RegisterPostHandler).Methods("POST")

	http.ListenAndServe(":8000", r)
}
