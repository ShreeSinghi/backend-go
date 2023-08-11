package api

import (
	"mvc/pkg/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()

	r.HandleFunc("/login", 			controllers.LoginHandler).Methods("GET")
	r.HandleFunc("/login", 			controllers.LoginPostHandler).Methods("POST")
	r.HandleFunc("/register", 		controllers.RegisterHandler).Methods("GET")
	r.HandleFunc("/register", 		controllers.Authenticate(controllers.RegisterPostHandler)).Methods("POST")
	r.HandleFunc("/home", 			controllers.Authenticate(controllers.HomeHandler)).Methods("GET")
	r.HandleFunc("/process-checks", controllers.Authenticate(controllers.ProcessChecks)).Methods("POST")


	http.ListenAndServe(":8000", r)
}
