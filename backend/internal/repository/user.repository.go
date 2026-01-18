package repository

import (
	"be-request-insident/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	// CreateUser(user *models.User) error
	// UpdateUser(user *models.User) error
	// DeleteUser(id string) error
}

type userRepository struct {
	DB *gorm.DB
	cache Cache
}

func NewUserRepository(db *gorm.DB, cache Cache) UserRepository {
	return &userRepository{DB: db, cache: cache}
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User

	cacheKey := "user:id:" + id

	if r.cache != nil {
		if s, err := r.cache.Get(ctx, cacheKey); err == nil {
			var cached models.User
			if jsonErr := json.Unmarshal([]byte(s), &cached); jsonErr == nil {
				fmt.Println("Cache hit for user ID:", id)
				return &cached, nil
			}
			_ = r.cache.Del(ctx, cacheKey)
		}
	}

	if err := r.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	if r.cache != nil {
		if data, err := json.Marshal(user); err == nil {
			_ = r.cache.Set(ctx, cacheKey, string(data),  5 * time.Minute)
		}
	}

	return &user, nil
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	cacheKey := "user:username:" + username

	if r.cache != nil {
		if s, err := r.cache.Get(ctx, cacheKey); err == nil {
			var cached models.User
			if jsonErr := json.Unmarshal([]byte(s), &cached); jsonErr == nil {
				return &cached, nil
			}
			_ = r.cache.Del(ctx, cacheKey)
		}
	}

	if err := r.DB.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	if r.cache != nil {
		if data, err := json.Marshal(user); err == nil {
			_ = r.cache.Set(ctx, cacheKey, string(data), 5 * time.Minute)
		}
	}

	
	return &user, nil
}