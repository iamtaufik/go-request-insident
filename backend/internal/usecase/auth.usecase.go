package usecase

import (
	"be-request-insident/internal/logger"
	"be-request-insident/internal/models"
	"be-request-insident/internal/repository"
	"be-request-insident/utility"
	"context"
	"errors"
	"log"
)


type AuthUsecase struct {
	userRepo repository.UserRepository
	appLog logger.Logger
}

func NewAuthUseCase(userRepo repository.UserRepository, appLog logger.Logger) *AuthUsecase {
	return &AuthUsecase{userRepo: userRepo, appLog: appLog}
}

func (u *AuthUsecase) Login(ctx context.Context, username, password string) (string, string, error) {
	isExist, err := u.userRepo.GetUserByUsername(username)

	if err != nil {
		return "", "", err
	}


	if isExist == nil {
		return "", "", errors.New("user not found")
	}

	isMatch := utility.CheckPasswordHash(password, isExist.Password)

	if !isMatch {
		log.Printf("Password does not match for user: %s", username)
		u.appLog.Error(ctx, "LOGIN_FAILED", "invalid password attempt", nil, map[string]interface{}{
			"username": username,
		})
		return "", "", errors.New("invalid credentials")
	}

	accessToken, err := utility.GenerateAccessToken(isExist.ID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utility.GenerateRefreshToken(isExist.ID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil 
}

func (u *AuthUsecase) Me(userID string) (*models.User, error) {
	user, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}