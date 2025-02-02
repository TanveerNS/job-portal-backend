package routes

import (
	"job-portal-backend/controllers"

	"github.com/gorilla/mux"
)

func InterviewScheduleRoutes(r *mux.Router) {
	r.HandleFunc("/interview-schedule", controllers.CreateInterviewSchedule).Methods("POST")
	r.HandleFunc("/interview-schedule", controllers.GetInterviewSchedule).Methods("GET")
	r.HandleFunc("/interview-schedule", controllers.UpdateInterviewSchedule).Methods("PUT")
	r.HandleFunc("/interview-schedule", controllers.DeleteInterviewSchedule).Methods("DELETE")
	r.HandleFunc("/interview-schedule/all", controllers.GetAllInterviewSchedules).Methods("GET")
}
