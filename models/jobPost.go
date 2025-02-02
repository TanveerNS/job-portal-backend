package models

import "time"

// JobPost represents a job posting for a company
type JobPost struct {
	ID                      string    `json:"id,omitempty" bson:"_id,omitempty"`
	JobTitle                string    `json:"job_title" bson:"job_title"`
	JobDes                  string    `json:"job_des" bson:"job_des"`
	FirstName               string    `json:"first_name" bson:"first_name"`
	LastName                string    `json:"last_name" bson:"last_name"`
	Email                   string    `json:"email" bson:"email"`
	JobType                 string    `json:"job_type" bson:"job_type"`
	Specialisms             string    `json:"specialisms" bson:"specialisms"`
	OfferedSalary           float64   `json:"offered_salary" bson:"offered_salary"`
	CareerLevel             string    `json:"career_level" bson:"career_level"`
	Experience              string    `json:"experience" bson:"experience"`
	Gender                  string    `json:"gender" bson:"gender"`
	Industry                string    `json:"industry" bson:"industry"`
	Qualification           string    `json:"qualification" bson:"qualification"`
	ApplicationDeadlineDate time.Time `json:"application_deadline_date" bson:"application_deadline_date"`
	Country                 string    `json:"country" bson:"country"`
	City                    string    `json:"city" bson:"city"`
	CompleteAddress         string    `json:"complete_address" bson:"complete_address"`
	Keywords                string    `json:"keywords" bson:"keywords"`
	CreatedAt               time.Time `json:"created_at" bson:"created_at"`
	Type                    string    `json:"type" bson:"type"`
	Status                  string    `json:"status" bson:"status"`
}
