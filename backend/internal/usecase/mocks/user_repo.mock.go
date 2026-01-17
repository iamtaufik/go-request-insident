package mocks

import (
	"be-request-insident/internal/models"
)

type UserRepoMock struct {
	GetUserByUsernameFn       func(username string) (*models.User, error)
	GetUserByIDFn             func(id string) (*models.User, error)
}

func (m *UserRepoMock) GetUserByUsername(username string) (*models.User, error) {
	return m.GetUserByUsernameFn(username)
}

func (m *UserRepoMock) GetUserByID(id string) (*models.User, error) {
	return m.GetUserByIDFn(id)
}