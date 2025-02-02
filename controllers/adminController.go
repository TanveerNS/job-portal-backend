package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"job-portal-backend/database"
	"job-portal-backend/helpers"
	"job-portal-backend/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func AdminRegister(w http.ResponseWriter, r *http.Request) {
	var admin models.Admin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		helpers.Respond(w, false, "Invalid input", nil, http.StatusBadRequest)
		return
	}

	// Check if the email is already registered
	collection := database.GetCollection("admins")
	var existingAdmin models.Admin
	err := collection.FindOne(context.Background(), bson.M{"email": admin.Email}).Decode(&existingAdmin)
	if err == nil {
		// Email already exists
		helpers.Respond(w, false, "Email is already registered", nil, http.StatusConflict)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.Respond(w, false, "Error hashing password", nil, http.StatusInternalServerError)
		return
	}
	admin.Password = string(hashedPassword)
	admin.CreatedAt = time.Now()

	// Insert the new admin
	_, err = collection.InsertOne(context.Background(), admin)
	if err != nil {
		helpers.Respond(w, false, "Error creating admin", nil, http.StatusInternalServerError)
		return
	}

	// Respond with status 201 Created after successful registration
	helpers.Respond(w, true, "Admin registered successfully", nil, http.StatusCreated)
}

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	var admin models.Admin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		helpers.Respond(w, false, "Invalid input", nil, http.StatusBadRequest)
		return
	}

	collection := database.GetCollection("admins")
	var dbAdmin models.Admin
	err := collection.FindOne(context.Background(), bson.M{"email": admin.Email}).Decode(&dbAdmin)
	if err != nil {
		helpers.Respond(w, false, "Admin not found", nil, http.StatusUnauthorized)
		return
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(dbAdmin.Password), []byte(admin.Password))
	if err != nil {
		helpers.Respond(w, false, "Invalid credentials", nil, http.StatusUnauthorized)
		return
	}

	// Create JWT
	token, err := helpers.GenerateJWT(dbAdmin.ID)
	if err != nil {
		helpers.Respond(w, false, "Error generating token", nil, http.StatusInternalServerError)
		return
	}

	// Prepare response data with token, userId, and email
	responseData := map[string]interface{}{
		"token":  token,
		"userId": dbAdmin.ID,
		"email":  dbAdmin.Email,
	}

	helpers.Respond(w, true, "Login successful", responseData, http.StatusOK)
}

func AdminAddUser(w http.ResponseWriter, r *http.Request) {
	var user models.Admin
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.Respond(w, false, "Invalid input", nil, http.StatusBadRequest)
		return
	}
	fmt.Print(user)
	if user.FirstName == "" || user.Email == "" || user.Mobile == "" || user.Password == "" {
		helpers.Respond(w, false, "First name, email, mobile, and password are required", nil, http.StatusBadRequest)
		return
	}

	// Check if the email already exists
	collection := database.GetCollection("admins")
	var existingUser models.Admin
	err := collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		helpers.Respond(w, false, "Email is already registered", nil, http.StatusConflict)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.Respond(w, false, "Error hashing password", nil, http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()

	// Insert the new user
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		helpers.Respond(w, false, "Error creating user", nil, http.StatusInternalServerError)
		return
	}

	// Respond with status 201 Created after successful user creation
	helpers.Respond(w, true, "User added successfully", nil, http.StatusCreated)
}

func AdminListUsers(w http.ResponseWriter, r *http.Request) {
	collection := database.GetCollection("admins")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		helpers.Respond(w, false, "Error fetching users", nil, http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var users []models.Admin
	if err := cursor.All(context.Background(), &users); err != nil {
		helpers.Respond(w, false, "Error decoding users", nil, http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, true, "Users fetched successfully", users, http.StatusOK)
}

func AdminGetUserDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	collection := database.GetCollection("users")
	var user models.User
	err := collection.FindOne(context.Background(), bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		helpers.Respond(w, false, "User not found", nil, http.StatusNotFound)
		return
	}

	helpers.Respond(w, true, "User details fetched successfully", user, http.StatusOK)
}

func AdminDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	collection := database.GetCollection("users")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": userId})
	if err != nil {
		helpers.Respond(w, false, "Error deleting user", nil, http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, true, "User deleted successfully", nil, http.StatusOK)
}
