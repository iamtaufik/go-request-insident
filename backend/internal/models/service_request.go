package models

import "time"

type ServiceRequest struct {
	ID                    string    `gorm:"column:id;primaryKey" json:"id"`
	RequestCode           string    `gorm:"column:request_code" json:"request_code"`
	RequestType           string    `gorm:"column:request_type;not null" json:"request_type"`
	Status                string    `gorm:"column:status;not null;default:DRAFT" json:"status"`
	ReportedFrom          string    `gorm:"column:reported_from" json:"reported_from"`
	Owner                 string    `gorm:"column:owner" json:"owner"`
	Company               string    `gorm:"column:company" json:"company"`
	Organization          string    `gorm:"column:organization" json:"organization"`
	Department            string    `gorm:"column:department" json:"department"`
	ProductName           string    `gorm:"column:product_name" json:"product_name"`
	ProductItem           string    `gorm:"column:product_item" json:"product_item"`
	ProductCategory       string    `gorm:"column:product_category" json:"product_category"`
	ProductParentCategory string    `gorm:"column:product_parent_category" json:"product_parent_category"`
	Symptom               string    `gorm:"column:symptom" json:"symptom"`
	RequestItem           string    `gorm:"column:request_item" json:"request_item"`
	Impact                string    `gorm:"column:impact" json:"impact"`
	Urgency               string    `gorm:"column:urgency" json:"urgency"`
	Priority              string    `gorm:"column:priority" json:"priority"`
	SolverGroup           string    `gorm:"column:solver_group" json:"solver_group"`
	Solver                string    `gorm:"column:solver" json:"solver"`
	CoordinatorGroup      string    `gorm:"column:coordinator_group" json:"coordinator_group"`
	Coordinator           string    `gorm:"column:coordinator" json:"coordinator"`
	SLAType               string    `gorm:"column:sla_type" json:"sla_type"`
	Summary               string    `gorm:"column:summary" json:"summary"`
	Note                  string    `gorm:"column:note" json:"note"`
	CreatedAt             time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt             time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}