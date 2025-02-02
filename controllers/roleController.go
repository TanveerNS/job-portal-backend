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

// Create Role
func CreateRole(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		helpers.Respond(w, false, "Invalid input", nil, http.StatusBadRequest)
		return
	}

	// Validate required fields
	if role.Name == "" {
		helpers.Respond(w, false, "Name is required", nil, http.StatusBadRequest)
		return
	}

	// Set created_at field to current time
	role.CreatedAt = time.Now()

	// Insert role into the collection
	collection := database.GetCollection("roles")
	_, err := collection.InsertOne(context.Background(), role)
	if err != nil {
		helpers.Respond(w, false, "Error creating role", nil, http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, true, "Role created successfully", nil, http.StatusCreated)
}

// Get Role by ID
func GetRole(w http.ResponseWriter, r *http.Request) {
	// Assuming the role ID is passed as a query parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		helpers.Respond(w, false, "ID is required", nil, http.StatusBadRequest)
		return
	}

	collection := database.GetCollection("roles")
	var role models.Role
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&role)
	if err != nil {
		helpers.Respond(w, false, "Role not found", nil, http.StatusNotFound)
		return
	}

	helpers.Respond(w, true, "Role found", role, http.StatusOK)
}

// Update Role by ID
func UpdateRole(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		helpers.Respond(w, false, "Invalid input", nil, http.StatusBadRequest)
		return
	}

	// Validate required fields for update
	if role.ID == "" || role.Name == "" {
		helpers.Respond(w, false, "ID and Name are required", nil, http.StatusBadRequest)
		return
	}

	// Update the role in the collection
	collection := database.GetCollection("roles")
	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": role.ID},
		bson.M{
			"$set": bson.M{
				"name":              role.Name,
				"company_apprvl":    role.CompanyApprvl,
				"company_edit":      role.CompanyEdit,
				"seeker_apprvl":     role.SeekerApprvl,
				"seeker_edit":       role.SeekerEdit,
				"recruiter_apprvl":  role.RecruiterApprvl,
				"recruiter_edit":    role.RecruiterEdit,
				"post_job":          role.PostJob,
				"data_download":     role.DataDownload,
				"del_job":           role.DelJob,
				"del_job_seeker":    role.DelJobSeeker,
				"del_recruiter":     role.DelRecruiter,
				"del_company":       role.DelCompany,
				"pass_reset":        role.PassReset,
				"chat_support":      role.ChatSupport,
				"subscription_chng": role.SubscriptionChng,
				"add_admin_hr":      role.AddAdminHr,
				"created_at":        role.CreatedAt,
			},
		},
	)
	if err != nil {
		helpers.Respond(w, false, "Error updating role", nil, http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, true, "Role updated successfully", nil, http.StatusOK)
}

// Delete Role by ID
func DeleteRole(w http.ResponseWriter, r *http.Request) {
	// Assuming the role ID is passed as a query parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		helpers.Respond(w, false, "ID is required", nil, http.StatusBadRequest)
		return
	}

	collection := database.GetCollection("roles")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		helpers.Respond(w, false, "Error deleting role", nil, http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, true, "Role deleted successfully", nil, http.StatusOK)
}

// Get All Roles
func GetAllRoles(w http.ResponseWriter, r *http.Request) {
	collection := database.GetCollection("roles")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		helpers.Respond(w, false, "Error retrieving roles", nil, http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var roles []models.Role
	for cursor.Next(context.Background()) {
		var role models.Role
		if err := cursor.Decode(&role); err != nil {
			helpers.Respond(w, false, "Error decoding role", nil, http.StatusInternalServerError)
			return
		}
		roles = append(roles, role)
	}

	helpers.Respond(w, true, "Roles retrieved successfully", roles, http.StatusOK)
}
