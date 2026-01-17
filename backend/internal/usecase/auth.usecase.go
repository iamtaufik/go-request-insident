package usecase

import (
	"be-request-insident/internal/models"
	"be-request-insident/internal/repository"
	"be-request-insident/utility"
	"errors"
	"log"
)


type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) Login(username, password string) (string, string, error) {
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

func (u *userUsecase) Me(userID string) (*models.User, error) {
	user, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}