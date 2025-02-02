package controllers

import (
	"context"
	"encoding/json"
	"job-portal-backend/database"
	"job-portal-backend/helpers"
	"job-portal-backend/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Create Interview Schedule
func CreateInterviewSched(w http.ResponseWriter, r *http.Request) {
	var sched models.AdminInterviewSched
	if err := json.NewDecoder(r.Body).Decode(&sched); err != nil {
		helpers.Respond(w, false, "Invalid input", nil, http.StatusBadRequest)
		return
	}

	// Validate required fields
	if sched.SeekerID == "" || sched.AdminID == "" || sched.HRName == "" || sched.HREmail == "" || sched.JobProfile == "" || sched.DateSchul.IsZero() || sched.FromTime == "" || sched.ToTime == "" {
		helpers.Respond(w, false, "All fields are required", nil, http.StatusBadRequest)
		return
	}

	sched.CreatedAt = time.Now()

	// Insert interview schedule into the collection
	collection := database.GetCollection("admin_interview_sched")
	_, err := collection.InsertOne(context.Background(), sched)
	if err != nil {
		helpers.Respond(w, false, "Error creating interview schedule", nil, http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, true, "Interview scheduled successfully", nil, http.StatusCreated)
}

// Get Interview Schedule by ID
func GetInterviewSched(w http.ResponseWriter, r *http.Request) {
	// Assuming the interview ID is passed as a query parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		helpers.Respond(w, false, "ID is required", nil, http.StatusBadRequest)
		return
	}

	collection := database.GetCollection("admin_interview_sched")
	var sched models.AdminInterviewSched
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&sched)
	if err != nil {
		helpers.Respond(w, false, "Interview not found", nil, http.StatusNotFound)
		return
	}

	helpers.Respond(w, true, "Interview schedule found", sched, http.StatusOK)
}

// Update Interview Schedule by ID
func UpdateInterviewSched(w http.ResponseWriter, r *http.Request) {
	var sched models.AdminInterviewSched
	if err := json.NewDecoder(r.Body).Decode(&sched); err != nil {
		helpers.Respond(w, false, "Invalid input", nil, http.StatusBadRequest)
		return
	}

	// Validate required fields for update
	if sched.ID == "" || sched.SeekerID == "" || sched.AdminID == "" || sched.HRName == "" || sched.HREmail == "" || sched.JobProfile == "" || sched.DateSchul.IsZero() || sched.FromTime == "" || sched.ToTime == "" {
		helpers.Respond(w, false, "All fields are required", nil, http.StatusBadRequest)
		return
	}

	// Update the interview schedule in the collection
	collection := database.GetCollection("admin_interview_sched")
	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": sched.ID},
		bson.M{
			"$set": bson.M{
				"seeker_id":   sched.SeekerID,
				"admin_id":    sched.AdminID,
				"hr_name":     sched.HRName,
				"hr_email":    sched.HREmail,
				"job_profile": sched.JobProfile,
				"description": sched.Description,
				"date_schul":  sched.DateSchul,
				"from_time":   sched.FromTime,
				"to_time":     sched.ToTime,
				"created_at":  sched.CreatedAt,
			},
		},
	)
	if err != nil {
		helpers.Respond(w, false, "Error updating interview schedule", nil, http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, true, "Interview schedule updated successfully", nil, http.StatusOK)
}

// Delete Interview Schedule by ID
func DeleteInterviewSched(w http.ResponseWriter, r *http.Request) {
	// Assuming the interview ID is passed as a query parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		helpers.Respond(w, false, "ID is required", nil, http.StatusBadRequest)
		return
	}

	collection := database.GetCollection("admin_interview_sched")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		helpers.Respond(w, false, "Error deleting interview schedule", nil, http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, true, "Interview schedule deleted successfully", nil, http.StatusOK)
}

// List all Interview Schedules
func GetAllInterviewScheds(w http.ResponseWriter, r *http.Request) {
	collection := database.GetCollection("admin_interview_sched")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		helpers.Respond(w, false, "Error retrieving interview schedules", nil, http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var schedules []models.AdminInterviewSched
	for cursor.Next(context.Background()) {
		var sched models.AdminInterviewSched
		if err := cursor.Decode(&sched); err != nil {
			helpers.Respond(w, false, "Error decoding interview schedule", nil, http.StatusInternalServerError)
			return
		}
		schedules = append(schedules, sched)
	}

	helpers.Respond(w, true, "Interview schedules retrieved successfully", schedules, http.StatusOK)
}
