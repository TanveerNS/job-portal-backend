package models

import "time"

type Admin struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	RoleID    string    `json:"role_id" bson:"role_id"`
	FirstName string    `json:"first_name" bson:"first_name"`
	LastName  string    `json:"last_name" bson:"last_name"`
	Email     string    `json:"email" bson:"email"`
	Mobile    string    `json:"mobile" bson:"mobile"`
	Password  string    `json:"password" bson:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
