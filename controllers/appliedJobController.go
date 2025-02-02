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

// Create Applied Job
func CreateAppliedJob(w http.ResponseWriter, r *http.Request) {
	var appliedJob models.AppliedJob
	if err := json.NewDecoder(r.Body).Decode(&appliedJob); err != nil {
		helpers.Respond(w, false, "Invalid input", nil, http.StatusBadRequest)
		return
	}

	// Validate required fields
	if appliedJob.SeekerID == "" || appliedJob.JobID == "" || appliedJob.Status == "" {
		helpers.Respond(w, false, "SeekerID, JobID, and Status are required", nil, http.StatusBadRequest)
		return
	}

	appliedJob.CreatedAt = time.Now()

	// Insert applied job into the collection
	collection := database.GetCollection("applied_job")
	_, err := collection.InsertOne(context.Background(), appliedJob)
	if err != nil {
		helpers.Respond(w, false, "Error applying for the job", nil, http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, true, "Job applied successfully", nil, http.StatusCreated)
}

// Get Applied Job by ID
func GetAppliedJob(w http.ResponseWriter, r *http.Request) {
	// Assuming the applied job ID is passed as a query parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		helpers.Respond(w, false, "ID is required", nil, http.StatusBadRequest)
		return
	}

	collection := database.GetCollection("applied_job")
	var appliedJob models.AppliedJob
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&appliedJob)
	if err != nil {
		helpers.Respond(w, false, "Applied job not found", nil, http.StatusNotFound)
		return
	}

	helpers.Respond(w, true, "Applied job found", appliedJob, http.StatusOK)
}

// Update Applied Job by ID
func UpdateAppliedJob(w http.ResponseWriter, r *http.Request) {
	var appliedJob models.AppliedJob
	if err := json.NewDecoder(r.Body).Decode(&appliedJob); err != nil {
		helpers.Respond(w, false, "Invalid input", nil, http.StatusBadRequest)
		return
	}

	// Validate required fields for update
	if appliedJob.ID == "" || appliedJob.SeekerID == "" || appliedJob.JobID == "" || appliedJob.Status == "" {
		helpers.Respond(w, false, "ID, SeekerID, JobID, and Status are required", nil, http.StatusBadRequest)
		return
	}

	// Update the applied job in the collection
	collection := database.GetCollection("applied_job")
	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": appliedJob.ID},
		bson.M{
			"$set": bson.M{
				"seeker_id":  appliedJob.SeekerID,
				"job_id":     appliedJob.JobID,
				"status":     appliedJob.Status,
				"created_at": appliedJob.CreatedAt,
			},
		},
	)
	if err != nil {
		helpers.Respond(w, false, "Error updating applied job", nil, http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, true, "Applied job updated successfully", nil, http.StatusOK)
}

// Delete Applied Job by ID
func DeleteAppliedJob(w http.ResponseWriter, r *http.Request) {
	// Assuming the applied job ID is passed as a query parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		helpers.Respond(w, false, "ID is required", nil, http.StatusBadRequest)
		return
	}

	collection := database.GetCollection("applied_job")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		helpers.Respond(w, false, "Error deleting applied job", nil, http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, true, "Applied job deleted successfully", nil, http.StatusOK)
}

// Get All Applied Jobs
func GetAllAppliedJobs(w http.ResponseWriter, r *http.Request) {
	collection := database.GetCollection("applied_job")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		helpers.Respond(w, false, "Error retrieving applied jobs", nil, http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var appliedJobs []models.AppliedJob
	for cursor.Next(context.Background()) {
		var appliedJob models.AppliedJob
		if err := cursor.Decode(&appliedJob); err != nil {
			helpers.Respond(w, false, "Error decoding applied job", nil, http.StatusInternalServerError)
			return
		}
		appliedJobs = append(appliedJobs, appliedJob)
	}

	helpers.Respond(w, true, "Applied jobs retrieved successfully", appliedJobs, http.StatusOK)
}
