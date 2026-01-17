package seed

import (
	"be-request-insident/internal/models"
	"be-request-insident/utility"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	users := []models.User{
		{
			ID:        uuid.NewString(),
			Username:  "taufikdev",
			Password:  utility.HashPassword("admin123"),
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uuid.NewString(),
			Username:  "user",
			Password:  utility.HashPassword("user123"),
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, user := range users {
		var exists models.User
		err := db.Where("username = ?", user.Username).First(&exists).Error

		if err == nil {
			log.Printf("[SEED] user %s already exists, skipping", user.Username)
			continue
		}

		if err := db.Create(&user).Error; err != nil {
			return err
		}

		log.Printf("[SEED] user %s created", user.Username)
	}

	return nil
}