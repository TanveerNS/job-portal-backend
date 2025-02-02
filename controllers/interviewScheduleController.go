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
func CreateInterviewSchedule(w http.ResponseWriter, r *http.Request) {
	var interviewSchedule models.InterviewSchedule
	if err := json.NewDecoder(r.Body).Decode(&interviewSchedule); err != nil {
		helpers.Respond(w, false, "Invalid input", nil, http.StatusBadRequest)
		return
	}

	// Validate required fields
	if interviewSchedule.SeekerID == "" || interviewSchedule.CompanyID == "" || interviewSchedule.HRName == "" || interviewSchedule.HREmail == "" {
		helpers.Respond(w, false, "SeekerID, CompanyID, HRName, and HREmail are required", nil, http.StatusBadRequest)
		return
	}

	// Set created_at field to the current time
	interviewSchedule.CreatedAt = time.Now()

	// Insert interview schedule into the collection
	collection := database.GetCollection("interview_schedule")
	_, err := collection.InsertOne(context.Background(), interviewSchedule)
	if err != nil {
		helpers.Respond(w, false, "Error creating interview schedule", nil, http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, true, "Interview schedule created successfully", nil, http.StatusCreated)
}

// Get Interview Schedule by ID
func GetInterviewSchedule(w http.ResponseWriter, r *http.Request) {
	// Assuming the interview schedule ID is passed as a query parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		helpers.Respond(w, false, "ID is required", nil, http.StatusBadRequest)
		return
	}

	collection := database.GetCollection("interview_schedule")
	var interviewSchedule models.InterviewSchedule
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&interviewSchedule)
	if err != nil {
		helpers.Respond(w, false, "Interview schedule not found", nil, http.StatusNotFound)
		return
	}

	helpers.Respond(w, true, "Interview schedule found", interviewSchedule, http.StatusOK)
}

// Update Interview Schedule by ID
func UpdateInterviewSchedule(w http.ResponseWriter, r *http.Request) {
	var interviewSchedule models.InterviewSchedule
	if err := json.NewDecoder(r.Body).Decode(&interviewSchedule); err != nil {
		helpers.Respond(w, false, "Invalid input", nil, http.StatusBadRequest)
		return
	}

	// Validate required fields for update
	if interviewSchedule.ID == "" || interviewSchedule.SeekerID == "" || interviewSchedule.CompanyID == "" || interviewSchedule.HRName == "" || interviewSchedule.HREmail == "" {
		helpers.Respond(w, false, "ID, SeekerID, CompanyID, HRName, and HREmail are required", nil, http.StatusBadRequest)
		return
	}

	// Update the interview schedule in the collection
	collection := database.GetCollection("interview_schedule")
	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": interviewSchedule.ID},
		bson.M{
			"$set": bson.M{
				"seeker_id":   interviewSchedule.SeekerID,
				"company_id":  interviewSchedule.CompanyID,
				"hr_name":     interviewSchedule.HRName,
				"hr_email":    interviewSchedule.HREmail,
				"job_profile": interviewSchedule.JobProfile,
				"description": interviewSchedule.Description,
				"date_schul":  interviewSchedule.DateSchul,
				"from_time":   interviewSchedule.FromTime,
				"to_time":     interviewSchedule.ToTime,
				"created_at":  interviewSchedule.CreatedAt,
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
func DeleteInterviewSchedule(w http.ResponseWriter, r *http.Request) {
	// Assuming the interview schedule ID is passed as a query parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		helpers.Respond(w, false, "ID is required", nil, http.StatusBadRequest)
		return
	}

	collection := database.GetCollection("interview_schedule")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		helpers.Respond(w, false, "Error deleting interview schedule", nil, http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, true, "Interview schedule deleted successfully", nil, http.StatusOK)
}

// Get All Interview Schedules
func GetAllInterviewSchedules(w http.ResponseWriter, r *http.Request) {
	collection := database.GetCollection("interview_schedule")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		helpers.Respond(w, false, "Error retrieving interview schedules", nil, http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var interviewSchedules []models.InterviewSchedule
	for cursor.Next(context.Background()) {
		var interviewSchedule models.InterviewSchedule
		if err := cursor.Decode(&interviewSchedule); err != nil {
			helpers.Respond(w, false, "Error decoding interview schedule", nil, http.StatusInternalServerError)
			return
		}
		interviewSchedules = append(interviewSchedules, interviewSchedule)
	}

	helpers.Respond(w, true, "Interview schedules retrieved successfully", interviewSchedules, http.StatusOK)
}
