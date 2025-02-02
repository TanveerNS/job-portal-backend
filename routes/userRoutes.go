package routes

import (
	"job-portal-backend/controllers"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	r.HandleFunc("/user/register", controllers.UserRegister).Methods("POST")
	r.HandleFunc("/user/login", controllers.UserLogin).Methods("POST")

	//add user
	r.HandleFunc("/user", controllers.AddUser).Methods("POST")
	r.HandleFunc("/users", controllers.ListUsers).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.GetUserDetails).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

}
