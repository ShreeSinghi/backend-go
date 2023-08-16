package api

import (
	"mvc/pkg/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Authenticate(controllers.DefaultHandler)).Methods("GET")
	r.HandleFunc("/register", controllers.ViewRegister).Methods("GET")
	r.HandleFunc("/login", controllers.ViewLogin).Methods("GET")
	r.HandleFunc("/logout", controllers.Authenticate(controllers.Logout)).Methods("GET")
	r.HandleFunc("/home", controllers.Authenticate(controllers.ViewHome)).Methods("GET")

	r.HandleFunc("/request-return", controllers.Authenticate(controllers.ViewRequestReturn)).Methods("GET")

	r.HandleFunc("/checkins", controllers.Authenticate(controllers.ViewCheckins)).Methods("GET")
	r.HandleFunc("/checkouts", controllers.Authenticate(controllers.ViewCheckouts)).Methods("GET")
	r.HandleFunc("/admin-requests", controllers.Authenticate(controllers.ViewAdminRequests)).Methods("GET")
	r.HandleFunc("/add-book", controllers.Authenticate(controllers.ViewAddBook)).Methods("GET")

	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/process-checks", controllers.Authenticate(controllers.ProcessChecks)).Methods("POST")
	r.HandleFunc("/add-book", controllers.Authenticate(controllers.AddBook)).Methods("POST")
	r.HandleFunc("/process-admin-requests", controllers.Authenticate(controllers.ProcessAdminRequests)).Methods("POST")
	r.HandleFunc("/request-admin", controllers.Authenticate(controllers.RequestAdmin)).Methods("POST")
	r.HandleFunc("/request-checkout", controllers.Authenticate(controllers.RequestCheckout)).Methods("POST")
	r.HandleFunc("/request-checkin", controllers.Authenticate(controllers.RequestCheckin)).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(controllers.NotFound)


	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.ListenAndServe(":8000", r)
}
