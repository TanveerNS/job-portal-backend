package routes

import (
	"job-portal-backend/controllers"

	"github.com/gorilla/mux"
)

func AppliedJobRoutes(r *mux.Router) {
	r.HandleFunc("/applied-job", controllers.CreateAppliedJob).Methods("POST")
	r.HandleFunc("/applied-job", controllers.GetAppliedJob).Methods("GET")
	r.HandleFunc("/applied-job", controllers.UpdateAppliedJob).Methods("PUT")
	r.HandleFunc("/applied-job", controllers.DeleteAppliedJob).Methods("DELETE")
	r.HandleFunc("/applied-job/all", controllers.GetAllAppliedJobs).Methods("GET")
}
