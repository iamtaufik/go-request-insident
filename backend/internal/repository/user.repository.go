package repository

import (
	"be-request-insident/internal/models"

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
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{DB: db}
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