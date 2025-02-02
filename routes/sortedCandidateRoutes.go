package routes

import (
	"job-portal-backend/controllers"

	"github.com/gorilla/mux"
)

func SortedCandidateRoutes(r *mux.Router) {
	r.HandleFunc("/sorted_candidates", controllers.CreateSortedCandidate).Methods("POST")
	r.HandleFunc("/sorted_candidates", controllers.GetSortedCandidates).Methods("GET")
	r.HandleFunc("/sorted_candidates/{id}", controllers.UpdateSortedCandidateStatus).Methods("PUT")
	r.HandleFunc("/sorted_candidates/{id}", controllers.DeleteSortedCandidate).Methods("DELETE")
}
