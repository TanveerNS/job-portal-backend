package models

import "time"

type User struct {
	ID                string    `json:"id,omitempty" bson:"_id,omitempty"`
	FullName          string    `json:"fullname" bson:"fullname"`
	Email             string    `json:"email" bson:"email"`
	Mobile            string    `json:"mobile" bson:"mobile"`
	Password          string    `json:"password" bson:"password"`
	Res               string    `json:"res,omitempty" bson:"res,omitempty"`
	Agree             bool      `json:"agree,omitempty" bson:"agree,omitempty"`
	Type              string    `json:"type,omitempty" bson:"type,omitempty"`
	MiddleName        string    `json:"middle_name,omitempty" bson:"middle_name,omitempty"`
	LastName          string    `json:"last_name,omitempty" bson:"last_name,omitempty"`
	ProfessionalTitle string    `json:"professional_title,omitempty" bson:"professional_title,omitempty"`
	Gender            string    `json:"gender,omitempty" bson:"gender,omitempty"`
	Dob               string    `json:"dob,omitempty" bson:"dob,omitempty"`
	Age               int       `json:"age,omitempty" bson:"age,omitempty"`
	CurrentIndustry   string    `json:"current_industry,omitempty" bson:"current_industry,omitempty"`
	Qualification     string    `json:"qualification,omitempty" bson:"qualification,omitempty"`
	ReadyToRelocate   bool      `json:"ready_to_relocate,omitempty" bson:"ready_to_relocate,omitempty"`
	TotalExperience   int       `json:"total_experience,omitempty" bson:"total_experience,omitempty"`
	CurrentSalary     float64   `json:"current_salary,omitempty" bson:"current_salary,omitempty"`
	ExpectedSalary    float64   `json:"expected_salary,omitempty" bson:"expected_salary,omitempty"`
	Keywords          string    `json:"keywords,omitempty" bson:"keywords,omitempty"`
	Bio               string    `json:"bio,omitempty" bson:"bio,omitempty"`
	Country           string    `json:"country,omitempty" bson:"country,omitempty"`
	State             string    `json:"state,omitempty" bson:"state,omitempty"`
	City              string    `json:"city,omitempty" bson:"city,omitempty"`
	Pincode           string    `json:"pincode,omitempty" bson:"pincode,omitempty"`
	FullAddress       string    `json:"full_address,omitempty" bson:"full_address,omitempty"`
	ProfilePic        string    `json:"profile_pic,omitempty" bson:"profile_pic,omitempty"`
	IsSubscribe       bool      `json:"is_subscribe,omitempty" bson:"is_subscribe,omitempty"`
	IsOtpVerify       bool      `json:"is_otp_verify,omitempty" bson:"is_otp_verify,omitempty"`
	IsActive          bool      `json:"is_active,omitempty" bson:"is_active,omitempty"`
	CreatedAt         time.Time `json:"created_at" bson:"created_at"`
}
