package controllers

import (
	"context"
	"encoding/json"
	"job-portal-backend/database"
	"job-portal-backend/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateSortedCandidate(w http.ResponseWriter, r *http.Request) {
	var sortedCandidate models.SortedCandidate
	if err := json.NewDecoder(r.Body).Decode(&sortedCandidate); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Set created_at timestamp
	sortedCandidate.CreatedAt = time.Now()

	collection := database.GetCollection("sorted_candidates")
	_, err := collection.InsertOne(context.Background(), sortedCandidate)
	if err != nil {
		http.Error(w, "Error creating sorted candidate", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sortedCandidate)
}

func GetSortedCandidates(w http.ResponseWriter, r *http.Request) {
	companyID := r.URL.Query().Get("company_id") // Optional filter for company

	filter := bson.M{}
	if companyID != "" {
		filter["company_id"] = companyID
	}

	collection := database.GetCollection("sorted_candidates")
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		http.Error(w, "Error retrieving sorted candidates", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var sortedCandidates []models.SortedCandidate
	for cursor.Next(context.Background()) {
		var sortedCandidate models.SortedCandidate
		if err := cursor.Decode(&sortedCandidate); err != nil {
			http.Error(w, "Error decoding sorted candidate", http.StatusInternalServerError)
			return
		}
		sortedCandidates = append(sortedCandidates, sortedCandidate)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sortedCandidates)
}

func UpdateSortedCandidateStatus(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Status string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Example: Update the status of a specific sorted candidate
	// Assume the sorted candidate ID is in the URL path
	id := r.URL.Query().Get("id")

	collection := database.GetCollection("sorted_candidates")
	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"status": input.Status}},
	)

	if err != nil {
		http.Error(w, "Error updating sorted candidate", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteSortedCandidate(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	collection := database.GetCollection("sorted_candidates")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		http.Error(w, "Error deleting sorted candidate", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
