package models

import "time"

type AdminInterviewSched struct {
	ID          string    `json:"id,omitempty" bson:"_id,omitempty"`
	SeekerID    string    `json:"seeker_id" bson:"seeker_id"`
	AdminID     string    `json:"admin_id" bson:"admin_id"`
	HRName      string    `json:"hr_name" bson:"hr_name"`
	HREmail     string    `json:"hr_email" bson:"hr_email"`
	JobProfile  string    `json:"job_profile" bson:"job_profile"`
	Description string    `json:"description" bson:"description"`
	DateSchul   time.Time `json:"date_schul" bson:"date_schul"`
	FromTime    string    `json:"from_time" bson:"from_time"`
	ToTime      string    `json:"to_time" bson:"to_time"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
