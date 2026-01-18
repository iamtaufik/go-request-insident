package usecase_test

import (
	"be-request-insident/internal/logger"
	"be-request-insident/internal/models"
	"be-request-insident/internal/usecase"
	"be-request-insident/internal/usecase/mocks"
	"be-request-insident/utility"
	"context"
	"testing"
)


func TestAuthUsecase_Login_Success(t *testing.T) {
	t.Setenv("JWT_SECRET_KEY", "testing-secret")
	t.Setenv("JWT_REFRESH_KEY", "testing-secret")

	hashed := utility.HashPassword("secret123")

	mockRepo := &mocks.UserRepoMock{
		GetUserByUsernameFn: func(ctx context.Context, username string) (*models.User, error) {
			if username != "test@mail.com" {
				t.Fatalf("unexpected username: %s", username)
			}
			return &models.User{
				ID:           "user-1",
				Username:        username,
				Password: string(hashed),
			}, nil
		},
	}

	mockUserSessionRepo := &mocks.UserSessionMock{
		GetSessionByUserIDFn: func(ctx context.Context, userID string) (*models.UserSession, error) {
			if userID != "user-1" {
				t.Fatalf("unexpected userID: %s", userID)
			}

			return nil, nil
		},
		CreateSessionFn: func(ctx context.Context, session *models.UserSession) error {
			return nil
		},

	}

	uc := usecase.NewAuthUseCase(mockRepo, mockUserSessionRepo, logger.NoopLogger{})

	access, refresh, err := uc.Login(t.Context(), "test@mail.com", "secret123")

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if access == "" || refresh == "" {
		t.Fatalf("expected tokens, got access=%q refresh=%q", access, refresh)
	}
}

func TestAuthUsecase_Login_InvalidPassword(t *testing.T) {
	t.Setenv("JWT_SECRET_KEY", "testing-secret")
	t.Setenv("JWT_REFRESH_KEY", "testing-secret")

	hashed := utility.HashPassword("rightpass")

	mockRepo := &mocks.UserRepoMock{
		GetUserByUsernameFn: func(ctx context.Context, username string) (*models.User, error) {
			if username != "test@mail.com" {
				t.Fatalf("unexpected username: %s", username)
			}
			return &models.User{
				ID:           "user-1",
				Username:        username,
				Password: string(hashed),
			}, nil
		},
	}

	mockUserSessionRepo := &mocks.UserSessionMock{
		GetSessionByUserIDFn: func(ctx context.Context, userID string) (*models.UserSession, error) {
			if userID != "user-1" {
				t.Fatalf("unexpected userID: %s", userID)
			}
		
			return nil, nil
		},
		CreateSessionFn: func(ctx context.Context, session *models.UserSession) error {
			return nil
		},
	}

	uc := usecase.NewAuthUseCase(mockRepo, mockUserSessionRepo, logger.NoopLogger{})

	_, _, err := uc.Login(t.Context(),"test@mail.com", "wrongpass")
	if err == nil || err.Error() != "invalid credentials" {
		t.Fatalf("expected invalid credentials, got %v", err)
	}
}

