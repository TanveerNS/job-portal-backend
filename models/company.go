package models

import "time"

type Company struct {
	ID          string    `json:"id,omitempty" bson:"_id,omitempty"`
	CompanyName string    `json:"company_name" bson:"company_name"`
	CompanyType string    `json:"company_type" bson:"company_type"`
	Designation string    `json:"designation,omitempty" bson:"designation,omitempty"`
	Email       string    `json:"email" bson:"email"`
	Website     string    `json:"website,omitempty" bson:"website,omitempty"`
	FoundedDate string    `json:"founded_date,omitempty" bson:"founded_date,omitempty"`
	Industry    string    `json:"industry,omitempty" bson:"industry,omitempty"`
	Bio         string    `json:"bio,omitempty" bson:"bio,omitempty"`
	Mobile      string    `json:"mobile" bson:"mobile"`
	Password    string    `json:"password" bson:"password"`
	Agree       bool      `json:"agree,omitempty" bson:"agree,omitempty"`
	Type        string    `json:"type,omitempty" bson:"type,omitempty"`
	Country     string    `json:"country,omitempty" bson:"country,omitempty"`
	State       string    `json:"state,omitempty" bson:"state,omitempty"`
	City        string    `json:"city,omitempty" bson:"city,omitempty"`
	PinCode     string    `json:"pin_code,omitempty" bson:"pin_code,omitempty"`
	Address     string    `json:"address,omitempty" bson:"address,omitempty"`
	Facebook    string    `json:"facebook,omitempty" bson:"facebook,omitempty"`
	Twitter     string    `json:"twitter,omitempty" bson:"twitter,omitempty"`
	Google      string    `json:"google,omitempty" bson:"google,omitempty"`
	Linkedin    string    `json:"linkedin,omitempty" bson:"linkedin,omitempty"`
	CompanyLogo string    `json:"company_logo,omitempty" bson:"company_logo,omitempty"`
	IsSubscribe bool      `json:"is_subscribe,omitempty" bson:"is_subscribe,omitempty"`
	IsOtpVerify bool      `json:"is_otp_verify,omitempty" bson:"is_otp_verify,omitempty"`
	IsActive    bool      `json:"is_active,omitempty" bson:"is_active,omitempty"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
