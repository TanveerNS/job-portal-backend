package routes

import (
	"job-portal-backend/controllers"

	"github.com/gorilla/mux"
)

func AdminRoutes(r *mux.Router) {
	r.HandleFunc("/admin/register", controllers.AdminRegister).Methods("POST")
	r.HandleFunc("/admin/login", controllers.AdminLogin).Methods("POST")

	// Admin managing users
	r.HandleFunc("/admin/user", controllers.AdminAddUser).Methods("POST")            // Add User
	r.HandleFunc("/admin/users", controllers.AdminListUsers).Methods("GET")          // List Users
	r.HandleFunc("/admin/user/{id}", controllers.AdminGetUserDetails).Methods("GET") // Get User Details
	r.HandleFunc("/admin/user/{id}", controllers.AdminDeleteUser).Methods("DELETE")  // Delete User
}
