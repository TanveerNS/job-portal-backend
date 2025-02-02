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
	"golang.org/x/crypto/bcrypt"
)

func CompanyRegister(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		helpers.Respond(w, false, "Invalid input", nil, http.StatusBadRequest)
		return
	}

	// Validate required fields
	if company.CompanyName == "" || company.CompanyType == "" || company.Email == "" || company.Mobile == "" || company.Password == "" {
		helpers.Respond(w, false, "Company name, company type, email, mobile, and password are required", nil, http.StatusBadRequest)
		return
	}

	// Check if the email already exists
	collection := database.GetCollection("companies")
	var existingCompany models.Company
	err := collection.FindOne(context.Background(), bson.M{"email": company.Email}).Decode(&existingCompany)
	if err == nil {
		helpers.Respond(w, false, "Email is already registered", nil, http.StatusConflict)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(company.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.Respond(w, false, "Error hashing password", nil, http.StatusInternalServerError)
		return
	}
	company.Password = string(hashedPassword)
	company.CreatedAt = time.Now()

	// Insert the new company
	_, err = collection.InsertOne(context.Background(), company)
	if err != nil {
		helpers.Respond(w, false, "Error creating company", nil, http.StatusInternalServerError)
		return
	}

	// Respond with status 201 Created after successful registration
	helpers.Respond(w, true, "Company registered successfully", nil, http.StatusCreated)
}

func CompanyLogin(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		helpers.Respond(w, false, "Invalid input", nil, http.StatusBadRequest)
		return
	}

	// Find company by email
	collection := database.GetCollection("companies")
	var dbCompany models.Company
	err := collection.FindOne(context.Background(), bson.M{"email": company.Email}).Decode(&dbCompany)
	if err != nil {
		helpers.Respond(w, false, "Company not found", nil, http.StatusUnauthorized)
		return
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(dbCompany.Password), []byte(company.Password))
	if err != nil {
		helpers.Respond(w, false, "Invalid credentials", nil, http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := helpers.GenerateJWT(dbCompany.ID)
	if err != nil {
		helpers.Respond(w, false, "Error generating token", nil, http.StatusInternalServerError)
		return
	}

	// Prepare response data with token, companyId, and email
	responseData := map[string]interface{}{
		"token":       token,
		"companyId":   dbCompany.ID,
		"email":       dbCompany.Email,
		"companyName": dbCompany.CompanyName,
	}

	helpers.Respond(w, true, "Login successful", responseData, http.StatusOK)
}
