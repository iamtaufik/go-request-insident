package models

import "time"

type ServiceRequestAttachment struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	ServiceRequestID string    `gorm:"column:service_request_id;not null" json:"service_request_id"`
	FileName         string    `gorm:"column:file_name;not null" json:"file_name"`
	FileURL          string    `gorm:"column:file_url;not null" json:"file_url"`
	MimeType         string    `gorm:"column:mime_type" json:"mime_type"`
	FileSize         int64     `gorm:"column:file_size" json:"file_size"`
	UploadedAt       time.Time `gorm:"column:uploaded_at;not null" json:"uploaded_at"`

	ServiceRequest   ServiceRequest `gorm:"foreignKey:ServiceRequestID;references:ID" json:"-"`
}