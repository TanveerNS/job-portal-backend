package models

import "time"

// Role represents a role with different permissions
type Role struct {
	ID               string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name             string    `json:"name" bson:"name"`
	CompanyApprvl    bool      `json:"company_apprvl" bson:"company_apprvl"`
	CompanyEdit      bool      `json:"company_edit" bson:"company_edit"`
	SeekerApprvl     bool      `json:"seeker_apprvl" bson:"seeker_apprvl"`
	SeekerEdit       bool      `json:"seeker_edit" bson:"seeker_edit"`
	RecruiterApprvl  bool      `json:"recruiter_apprvl" bson:"recruiter_apprvl"`
	RecruiterEdit    bool      `json:"recruiter_edit" bson:"recruiter_edit"`
	PostJob          bool      `json:"post_job" bson:"post_job"`
	DataDownload     bool      `json:"data_download" bson:"data_download"`
	DelJob           bool      `json:"del_job" bson:"del_job"`
	DelJobSeeker     bool      `json:"del_job_seeker" bson:"del_job_seeker"`
	DelRecruiter     bool      `json:"del_recruiter" bson:"del_recruiter"`
	DelCompany       bool      `json:"del_company" bson:"del_company"`
	PassReset        bool      `json:"pass_reset" bson:"pass_reset"`
	ChatSupport      bool      `json:"chat_support" bson:"chat_support"`
	SubscriptionChng bool      `json:"subscription_chng" bson:"subscription_chng"`
	AddAdminHr       bool      `json:"add_admin_hr" bson:"add_admin_hr"`
	CreatedAt        time.Time `json:"created_at" bson:"created_at"`
}
