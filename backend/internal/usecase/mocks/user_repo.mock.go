package mocks

import (
	"be-request-insident/internal/models"
	"context"
)

type UserRepoMock struct {
	GetUserByUsernameFn       func(ctx context.Context, username string) (*models.User, error)
	GetUserByIDFn             func(ctx context.Context, id string) (*models.User, error)
}

func (m *UserRepoMock) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return m.GetUserByUsernameFn(ctx, username)
}

func (m *UserRepoMock) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return m.GetUserByIDFn(ctx, id)
}