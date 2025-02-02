package routes

import (
	"job-portal-backend/controllers"

	"github.com/gorilla/mux"
)

func CompanyRoutes(r *mux.Router) {
	r.HandleFunc("/company/register", controllers.CompanyRegister).Methods("POST")
	r.HandleFunc("/company/login", controllers.CompanyLogin).Methods("POST")
}
