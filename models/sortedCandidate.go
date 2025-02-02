package models

import "time"

// SortedCandidate represents a candidate that's been matched or sorted based on some criteria
type SortedCandidate struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"` // Unique identifier
	SeekerID  string    `json:"seeker_id" bson:"seeker_id"`        // Refers to the job seeker's user ID
	CompanyID string    `json:"company_id" bson:"company_id"`      // Refers to the company ID (job posting entity)
	Status    string    `json:"status" bson:"status"`              // Status of the match (e.g., "applied", "interviewed", "hired")
	CreatedAt time.Time `json:"created_at" bson:"created_at"`      // Date the candidate was matched or sorted
}
