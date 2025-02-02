package controllers

import (
	"context"
	"encoding/json"
	"job-portal-backend/database"
	"job-portal-backend/helpers"
	"job-portal-backend/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.Respond(w, false, "Invalid input", nil, http.StatusBadRequest)
		return
	}

	// Validate required fields
	if user.FullName == "" || user.Email == "" || user.Mobile == "" || user.Password == "" {
		helpers.Respond(w, false, "Full name, email, mobile, and password are required", nil, http.StatusBadRequest)
		return
	}

	// Check if the email already exists
	collection := database.GetCollection("users")
	var existingUser models.User
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

	// Respond with status 201 Created after successful registration
	helpers.Respond(w, true, "User registered successfully", nil, http.StatusCreated)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.Respond(w, false, "Invalid input", nil, http.StatusBadRequest)
		return
	}

	// Find user by email
	collection := database.GetCollection("users")
	var dbUser models.User
	err := collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&dbUser)
	if err != nil {
		helpers.Respond(w, false, "User not found", nil, http.StatusUnauthorized)
		return
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		helpers.Respond(w, false, "Invalid credentials", nil, http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := helpers.GenerateJWT(dbUser.ID)
	if err != nil {
		helpers.Respond(w, false, "Error generating token", nil, http.StatusInternalServerError)
		return
	}

	// Prepare response data with token, userId, and email
	responseData := map[string]interface{}{
		"token":  token,
		"userId": dbUser.ID,
		"email":  dbUser.Email,
	}

	helpers.Respond(w, true, "Login successful", responseData, http.StatusOK)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.Respond(w, false, "Invalid input", nil, http.StatusBadRequest)
		return
	}

	// Validate required fields
	if user.FullName == "" || user.Email == "" || user.Mobile == "" || user.Password == "" {
		helpers.Respond(w, false, "Full name, email, mobile, and password are required", nil, http.StatusBadRequest)
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
	collection := database.GetCollection("users")
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		helpers.Respond(w, false, "Error creating user", nil, http.StatusInternalServerError)
		return
	}

	// Respond with status 201 Created after successful registration
	helpers.Respond(w, true, "User added successfully", nil, http.StatusCreated)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	collection := database.GetCollection("users")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		helpers.Respond(w, false, "Error fetching users", nil, http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var users []models.User
	if err := cursor.All(context.Background(), &users); err != nil {
		helpers.Respond(w, false, "Error decoding users", nil, http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, true, "Users fetched successfully", users, http.StatusOK)
}

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
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
