package repository

import (
	"be-request-insident/internal/models"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(id string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	// CreateUser(user *models.User) error
	// UpdateUser(user *models.User) error
	// DeleteUser(id string) error
}

type userRepository struct {
	DB *gorm.DB
	redis *redis.Client
}

func NewUserRepository(db *gorm.DB, redis *redis.Client) UserRepository {
	return &userRepository{DB: db, redis: redis}
}

func (r *userRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return &user, nil
}