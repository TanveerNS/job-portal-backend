package routes

import (
	"job-portal-backend/controllers"

	"github.com/gorilla/mux"
)

func RoleRoutes(r *mux.Router) {
	r.HandleFunc("/role", controllers.CreateRole).Methods("POST")
	r.HandleFunc("/role", controllers.GetRole).Methods("GET")
	r.HandleFunc("/role", controllers.UpdateRole).Methods("PUT")
	r.HandleFunc("/role", controllers.DeleteRole).Methods("DELETE")
	r.HandleFunc("/role/all", controllers.GetAllRoles).Methods("GET")
}
