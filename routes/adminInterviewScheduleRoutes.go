package routes

import (
	"job-portal-backend/controllers"

	"github.com/gorilla/mux"
)

func AdminInterviewScheduleRoutes(r *mux.Router) {
	r.HandleFunc("/admin/interview-sched", controllers.CreateInterviewSched).Methods("POST")
	r.HandleFunc("/admin/interview-sched", controllers.GetInterviewSched).Methods("GET")
	r.HandleFunc("/admin/interview-sched", controllers.UpdateInterviewSched).Methods("PUT")
	r.HandleFunc("/admin/interview-sched", controllers.DeleteInterviewSched).Methods("DELETE")
	r.HandleFunc("/admin/interview-sched/all", controllers.GetAllInterviewScheds).Methods("GET")
}
