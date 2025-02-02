package models

import "time"

// AppliedJob represents the job application made by a job seeker.
type AppliedJob struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	SeekerID  string    `json:"seeker_id" bson:"seeker_id"`
	JobID     string    `json:"job_id" bson:"job_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Status    string    `json:"status" bson:"status"` // "applied", "interview", "hired", "rejected"
}
