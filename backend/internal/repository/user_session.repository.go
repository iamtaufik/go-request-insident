package repository

import (
	"be-request-insident/internal/models"
	"context"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type UserSessionRepository interface {
	GetSessionByUserID(ctx context.Context, userID string) (*models.UserSession, error)
	CreateSession(ctx context.Context, session *models.UserSession) error
	UpdateSession(ctx context.Context, session *models.UserSession) error
	UpdateSessionByUserID(ctx context.Context, userID string, session *models.UserSession) error
	DeleteSession(ctx context.Context, id string) error
}

type userSessionRepository struct {
	DB    *gorm.DB
	cache Cache
}

func NewUserSessionRepository(db *gorm.DB, cache Cache) UserSessionRepository {
	return &userSessionRepository{DB: db, cache: cache}
}

func (r *userSessionRepository) GetSessionByUserID(ctx context.Context, userID string) (*models.UserSession, error) {
	var session models.UserSession
	cacheKey := "user_session:user_id:" + userID

	if r.cache != nil {
		if s, err := r.cache.Get(ctx, cacheKey); err == nil {
			var cached models.UserSession
			if jsonErr := json.Unmarshal([]byte(s), &cached); jsonErr == nil {
				return &cached, nil
			}
			_ = r.cache.Del(ctx, cacheKey)
		}
	}

	if err := r.DB.First(&session, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	if r.cache != nil {
		if data, err := json.Marshal(session); err == nil {
			_ = r.cache.Set(ctx, cacheKey, string(data), 5*time.Minute)
		}
	}
	return &session, nil
}

func (r *userSessionRepository) CreateSession(ctx context.Context, session *models.UserSession) error {
	if err := r.DB.Create(session).Error; err != nil {
		return err
	}
	return nil
}

func (r *userSessionRepository) UpdateSession(ctx context.Context, session *models.UserSession) error {
	cacheKey := "user_session:user_id:" + session.UserID
	if r.cache != nil {
		_ = r.cache.Del(ctx, cacheKey)
	}
	if err := r.DB.Save(session).Error; err != nil {
		return err
	}

	return nil
}

func (r *userSessionRepository) UpdateSessionByUserID(ctx context.Context, userID string, session *models.UserSession) error {
	cacheKey := "user_session:user_id:" + userID
	if r.cache != nil {
		_ = r.cache.Del(ctx, cacheKey)
	}
	if err := r.DB.Model(&models.UserSession{}).Where("user_id = ?", userID).Updates(session).Error; err != nil {
		return err
	}

	return nil
}

func (r *userSessionRepository) DeleteSession(ctx context.Context, id string) error {
	if err := r.DB.Delete(&models.UserSession{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
