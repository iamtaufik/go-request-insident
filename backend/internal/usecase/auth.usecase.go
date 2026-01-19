package usecase

import (
	"be-request-insident/internal/logger"
	"be-request-insident/internal/models"
	"be-request-insident/internal/repository"
	"be-request-insident/utility"
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

type AuthUsecase struct {
	userRepo    repository.UserRepository
	sessionRepo repository.UserSessionRepository
	appLog      logger.Logger
}

func NewAuthUseCase(userRepo repository.UserRepository, sessionRepo repository.UserSessionRepository, appLog logger.Logger) *AuthUsecase {
	return &AuthUsecase{userRepo: userRepo, sessionRepo: sessionRepo, appLog: appLog}
}

func (u *AuthUsecase) Login(ctx context.Context, username, password string) (string, string, error) {
	user, err := u.userRepo.GetUserByUsername(ctx, username)

	if err != nil {
		return "", "", err
	}

	if user == nil {
		return "", "", errors.New("user not found")
	}

	session, _ := u.sessionRepo.GetSessionByUserID(ctx, user.ID)

	isMatch := utility.CheckPasswordHash(password, user.Password)

	if !isMatch {
		log.Printf("Password does not match for user: %s", username)
		u.appLog.Error(ctx, "LOGIN_FAILED", "invalid password attempt", nil, map[string]interface{}{
			"username": username,
		})
		return "", "", errors.New("invalid credentials")
	}

	accessToken, err := utility.GenerateAccessToken(user.ID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utility.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", err
	}

	if session != nil && session.Status == "ACTIVE" {
		log.Printf("User %s already has an active session", username)
		return "", "", errors.New("user already has an active session")
	}

	sessionID := uuid.New().String()

	refreshExp := time.Now().Add(7 * 24 * time.Hour)

	if session == nil {
		err = u.sessionRepo.CreateSession(ctx, &models.UserSession{
			ID:               uuid.New().String(),
			UserID:           user.ID,
			SessionID:        sessionID,
			RefreshTokenHash: refreshToken,
			Status:           "ACTIVE",
			RefreshExpiresAt: refreshExp,
			LastSeenAt:       time.Now(),
			RevokedAt:        nil,
			ExpiresAt:        refreshExp,
		})
		if err != nil {
			return "", "", err
		}

	} else {
		err = u.sessionRepo.UpdateSessionByUserID(ctx, user.ID, &models.UserSession{
			SessionID:        sessionID,
			RefreshTokenHash: refreshToken,
			Status:           "ACTIVE",
			RefreshExpiresAt: refreshExp,
			LastSeenAt:       time.Now(),
			RevokedAt:        nil,
			ExpiresAt:        refreshExp,
		})
		if err != nil {
			return "", "", err
		}
	}

	return accessToken, refreshToken, nil
}

func (u *AuthUsecase) Me(ctx context.Context, userID string) (*models.User, error) {
	user, err := u.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (u *AuthUsecase) Logout(ctx context.Context, userID string) error {
	session, err := u.sessionRepo.GetSessionByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if session == nil || session.Status != "ACTIVE" {
		return errors.New("no active session found")
	}

	now := time.Now()
	session.Status = "INACTIVE"
	session.RevokedAt = &now

	err = u.sessionRepo.UpdateSession(ctx, session)
	if err != nil {
		return err
	}

	u.appLog.Log(ctx, logger.AppLog{
		TS:      time.Now(),
		Event:   "LOGOUT_SUCCESS",
		Level:   logger.INFO,
		Service: "auth",
		Message: "user logged out successfully",
		UserID:  &userID,
	})

	return nil
}

func (u *AuthUsecase) RefreshToken(ctx context.Context, userID, refreshToken string) (string, string, error) {
	session, err := u.sessionRepo.GetSessionByUserID(ctx, userID)
	if err != nil {
		return "", "", err
	}

	if session == nil || session.Status != "ACTIVE" {
		return "", "", errors.New("no active session found")
	}

	if session.RefreshTokenHash != refreshToken {
		return "", "", errors.New("invalid refresh token")
	}

	if time.Now().After(session.RefreshExpiresAt) {
		return "", "", errors.New("refresh token expired")
	}

	newAccessToken, err := utility.GenerateAccessToken(userID)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := utility.GenerateRefreshToken(userID)
	if err != nil {
		return "", "", err
	}

	session.RefreshTokenHash = newRefreshToken
	session.RefreshExpiresAt = time.Now().Add(7 * 24 * time.Hour)
	session.LastSeenAt = time.Now()

	err = u.sessionRepo.UpdateSession(ctx, session)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}
