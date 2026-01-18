package repository

import (
	"be-request-insident/internal/models"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserSessionRepository interface {
	GetSessionByUserID(userID string) (*models.UserSession, error)
	CreateSession(session *models.UserSession) error
	UpdateSession(session *models.UserSession) error
	UpdateSessionByUserID(userID string, session *models.UserSession) error
	DeleteSession(id string) error
}

type userSessionRepository struct {
	DB *gorm.DB
	redis *redis.Client
}

func NewUserSessionRepository(db *gorm.DB, redis *redis.Client) UserSessionRepository {
	return &userSessionRepository{DB: db, redis: redis}
}


func (r *userSessionRepository) GetSessionByUserID(userID string) (*models.UserSession, error) {
	var session models.UserSession
	if err := r.DB.First(&session, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *userSessionRepository) CreateSession(session *models.UserSession) error {
	if err := r.DB.Create(session).Error; err != nil {
		return err
	}
	return nil
}

func (r *userSessionRepository) UpdateSession(session *models.UserSession) error {
	if err := r.DB.Save(session).Error; err != nil {
		return err
	}
	return nil
}

func (r *userSessionRepository) UpdateSessionByUserID(userID string, session *models.UserSession) error {
	if err := r.DB.Model(&models.UserSession{}).Where("user_id = ?", userID).Updates(session).Error; err != nil {
		return err
	}

	return nil
}

func (r *userSessionRepository) DeleteSession(id string) error {
	if err := r.DB.Delete(&models.UserSession{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}