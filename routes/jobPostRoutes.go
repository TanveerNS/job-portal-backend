package routes

import (
	"job-portal-backend/controllers"

	"github.com/gorilla/mux"
)

func JobPostRoutes(r *mux.Router) {
	r.HandleFunc("/job-post", controllers.CreateJobPost).Methods("POST")
	r.HandleFunc("/job-post", controllers.GetJobPost).Methods("GET")
	r.HandleFunc("/job-post", controllers.UpdateJobPost).Methods("PUT")
	r.HandleFunc("/job-post", controllers.DeleteJobPost).Methods("DELETE")
	r.HandleFunc("/job-post/all", controllers.GetAllJobPosts).Methods("GET")
}
