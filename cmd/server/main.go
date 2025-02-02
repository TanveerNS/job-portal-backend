package main

import (
	"job-portal-backend/database"
	"job-portal-backend/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Ensure Connect is called first
	log.Println("Connecting to MongoDB...")
	database.Connect() // This should log "MongoDB connected" if successful

	// Set up routing
	router := mux.NewRouter()

	// Define routes

	routes.AdminRoutes(router)
	routes.UserRoutes(router)
	routes.CompanyRoutes(router)
	routes.AdminInterviewScheduleRoutes(router)
	routes.AppliedJobRoutes(router)
	routes.InterviewScheduleRoutes(router)
	routes.JobPostRoutes(router)
	routes.RoleRoutes(router)
	routes.SortedCandidateRoutes(router)

	// Enable CORS with custom settings
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5173", "http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With", "Accept"},
		AllowCredentials: true,
	}

	// Apply the CORS handler with these options
	handler := cors.New(corsOptions).Handler(router)

	// Start the server
	log.Println("Server started at http://localhost:8000")
	if err := http.ListenAndServe(":8000", handler); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
