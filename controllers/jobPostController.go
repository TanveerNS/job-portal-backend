package controllers

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"job-portal-backend/database"
	"job-portal-backend/helpers"
	"job-portal-backend/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Create Job Post
func CreateJobPost(w http.ResponseWriter, r *http.Request) {
	var jobPost models.JobPost
	if err := json.NewDecoder(r.Body).Decode(&jobPost); err != nil {
		helpers.Respond(w, false, "Invalid input", nil, http.StatusBadRequest)
		return
	}

	// Validate required fields
	if jobPost.JobTitle == "" || jobPost.Email == "" || jobPost.JobDes == "" || jobPost.Country == "" || jobPost.City == "" {
		helpers.Respond(w, false, "JobTitle, Email, JobDes, Country, and City are required", nil, http.StatusBadRequest)
		return
	}

	// Set created_at field to the current time
	jobPost.CreatedAt = time.Now()

	// Insert job post into the collection
	collection := database.GetCollection("job_post")
	_, err := collection.InsertOne(context.Background(), jobPost)
	if err != nil {
		helpers.Respond(w, false, "Error creating job post", nil, http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, true, "Job post created successfully", nil, http.StatusCreated)
}

func GetJobPost(w http.ResponseWriter, r *http.Request) {
	// Assuming the job post ID is passed as a query parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		helpers.Respond(w, false, "ID is required", nil, http.StatusBadRequest)
		return
	}

	// Try to convert the string ID into an ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helpers.Respond(w, false, "Invalid ID format", nil, http.StatusBadRequest)
		return
	}

	// Proceed to query with the ObjectID
	collection := database.GetCollection("job_post")
	var jobPost models.JobPost
	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&jobPost)
	if err != nil {
		helpers.Respond(w, false, "Job post not found", nil, http.StatusNotFound)
		return
	}

	helpers.Respond(w, true, "Job post found", jobPost, http.StatusOK)
}

// Update Job Post by ID
func UpdateJobPost(w http.ResponseWriter, r *http.Request) {
	var jobPost models.JobPost
	if err := json.NewDecoder(r.Body).Decode(&jobPost); err != nil {
		helpers.Respond(w, false, "Invalid input", nil, http.StatusBadRequest)
		return
	}

	// Validate required fields for update
	if jobPost.ID == "" || jobPost.JobTitle == "" || jobPost.Email == "" || jobPost.JobDes == "" || jobPost.Country == "" || jobPost.City == "" {
		helpers.Respond(w, false, "ID, JobTitle, Email, JobDes, Country, and City are required", nil, http.StatusBadRequest)
		return
	}

	// Update the job post in the collection
	collection := database.GetCollection("job_post")
	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": jobPost.ID},
		bson.M{
			"$set": bson.M{
				"job_title":                 jobPost.JobTitle,
				"job_des":                   jobPost.JobDes,
				"first_name":                jobPost.FirstName,
				"last_name":                 jobPost.LastName,
				"email":                     jobPost.Email,
				"job_type":                  jobPost.JobType,
				"specialisms":               jobPost.Specialisms,
				"offered_salary":            jobPost.OfferedSalary,
				"career_level":              jobPost.CareerLevel,
				"experience":                jobPost.Experience,
				"gender":                    jobPost.Gender,
				"industry":                  jobPost.Industry,
				"qualification":             jobPost.Qualification,
				"application_deadline_date": jobPost.ApplicationDeadlineDate,
				"country":                   jobPost.Country,
				"city":                      jobPost.City,
				"complete_address":          jobPost.CompleteAddress,
				"keywords":                  jobPost.Keywords,
				"created_at":                jobPost.CreatedAt,
				"type":                      jobPost.Type,
				"status":                    jobPost.Status,
			},
		},
	)
	if err != nil {
		helpers.Respond(w, false, "Error updating job post", nil, http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, true, "Job post updated successfully", nil, http.StatusOK)
}

// Delete Job Post by ID
func DeleteJobPost(w http.ResponseWriter, r *http.Request) {
	// Assuming the job post ID is passed as a query parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		helpers.Respond(w, false, "ID is required", nil, http.StatusBadRequest)
		return
	}

	collection := database.GetCollection("job_post")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		helpers.Respond(w, false, "Error deleting job post", nil, http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, true, "Job post deleted successfully", nil, http.StatusOK)
}

// Get All Job Posts
func GetAllJobPosts(w http.ResponseWriter, r *http.Request) {
	collection := database.GetCollection("job_post")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		helpers.Respond(w, false, "Error retrieving job posts", nil, http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var jobPosts []models.JobPost
	for cursor.Next(context.Background()) {
		var jobPost models.JobPost
		if err := cursor.Decode(&jobPost); err != nil {
			helpers.Respond(w, false, "Error decoding job post", nil, http.StatusInternalServerError)
			return
		}
		jobPosts = append(jobPosts, jobPost)
	}

	helpers.Respond(w, true, "Job posts retrieved successfully", jobPosts, http.StatusOK)
}
