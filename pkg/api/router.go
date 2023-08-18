package api

import (
	"mvc/pkg/controllers"
	"mvc/pkg/middleware"

	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()

	// common get requests
	r.HandleFunc("/", middleware.Authenticate(controllers.DefaultHandler)).Methods("GET")
	r.HandleFunc("/register", controllers.ViewRegister).Methods("GET")
	r.HandleFunc("/login", controllers.ViewLogin).Methods("GET")
	r.HandleFunc("/logout", middleware.Authenticate(controllers.Logout)).Methods("GET")
	r.HandleFunc("/home", middleware.Authenticate(controllers.ViewHome)).Methods("GET")

	// user get requests
	r.HandleFunc("/request-return", middleware.Authenticate(controllers.ViewRequestReturn)).Methods("GET")

	// admin get requests
	r.HandleFunc("/checkins", middleware.Authenticate(controllers.ViewCheckins)).Methods("GET")
	r.HandleFunc("/checkouts", middleware.Authenticate(controllers.ViewCheckouts)).Methods("GET")
	r.HandleFunc("/admin-requests", middleware.Authenticate(controllers.ViewAdminRequests)).Methods("GET")
	r.HandleFunc("/add-book", middleware.Authenticate(controllers.ViewAddBook)).Methods("GET")

	// post requests
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/process-checks", middleware.Authenticate(controllers.ProcessChecks)).Methods("POST")
	r.HandleFunc("/add-book", middleware.Authenticate(controllers.AddBook)).Methods("POST")
	r.HandleFunc("/process-admin-requests", middleware.Authenticate(controllers.ProcessAdminRequests)).Methods("POST")
	r.HandleFunc("/request-checkout", middleware.Authenticate(controllers.RequestCheckout)).Methods("POST")
	r.HandleFunc("/request-checkin", middleware.Authenticate(controllers.RequestCheckin)).Methods("POST")
	r.HandleFunc("/request-admin", middleware.Authenticate(controllers.RequestAdmin)).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(controllers.NotFound)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.ListenAndServe(":8000", r)
}
