package models

import "time"

type UserSession struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	UserID           string    `gorm:"column:user_id;not null" json:"user_id"`
	SessionID        string    `gorm:"column:session_id;not null" json:"session_id"`
	Status           string    `gorm:"column:status;not null" json:"status"`
	RefreshTokenHash string    `gorm:"column:refresh_token_hash;not null" json:"refresh_token_hash"`
	RefreshExpiresAt time.Time `gorm:"column:refresh_expires_at;not null" json:"refresh_expires_at"`
	IPAddress        string    `gorm:"column:ip_address" json:"ip_address"`
	UserAgent        string    `gorm:"column:user_agent" json:"user_agent"`
	CreatedAt        time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	LastSeenAt       time.Time `gorm:"column:last_seen_at" json:"last_seen_at"`
	RevokedAt        *time.Time `gorm:"column:revoked_at" json:"revoked_at"`
	ExpiresAt        time.Time `gorm:"column:expires_at" json:"expires_at"`

	User             User      `gorm:"foreignKey:UserID;references:ID"`
}