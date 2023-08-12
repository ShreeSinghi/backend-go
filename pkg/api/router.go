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
	r.HandleFunc("/register", controllers.RegisterHandler).Methods("GET")
	r.HandleFunc("/register", controllers.RegisterPostHandler).Methods("POST")
	r.HandleFunc("/home", controllers.Authenticate(controllers.HomeHandler)).Methods("GET")
	r.HandleFunc("/process-checks", controllers.Authenticate(controllers.ProcessChecks)).Methods("POST")

	r.HandleFunc("/add-book", controllers.Authenticate(controllers.AddBook)).Methods("POST")
	r.HandleFunc("/process-admin-requests", controllers.Authenticate(controllers.ProcessAdminRequests)).Methods("POST")

	r.HandleFunc("/request-admin", controllers.Authenticate(controllers.RequestAdmin)).Methods("POST")
	r.HandleFunc("/return-book", controllers.Authenticate(controllers.ReturnBook)).Methods("POST")
	r.HandleFunc("/request-checkout", controllers.Authenticate(controllers.RequestCheckout)).Methods("POST")

	http.ListenAndServe(":8000", r)
}
