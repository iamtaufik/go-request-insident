package models

import (
	"time"

	"gorm.io/gorm"
)

type RUserServiceRequest struct {
	ID               string         `gorm:"column:id;primaryKey" json:"id"`
	UserID           string         `gorm:"column:user_id;not null" json:"user_id"`
	ServiceRequestID string         `gorm:"column:service_request_id;not null" json:"service_request_id"`
	IsActive         bool           `gorm:"column:is_active;default:1" json:"is_active"`
	CreatedAt        time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`

	User 			User 			`gorm:"foreignKey:UserID;references:ID"`
	ServiceRequest 	ServiceRequest 	`gorm:"foreignKey:ServiceRequestID;references:ID"`
}