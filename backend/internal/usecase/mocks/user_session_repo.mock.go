package mocks

import (
	"be-request-insident/internal/models"
	"context"
)

type UserSessionMock struct {
	GetSessionByUserIDFn func(ctx context.Context, userID string) (*models.UserSession, error)
	CreateSessionFn    func(ctx context.Context, session *models.UserSession) error
	UpdateSessionFn    func(ctx context.Context, session *models.UserSession) error
	UpdateSessionByUserIDFn func(ctx context.Context, userID string, session *models.UserSession) error
	DeleteSessionFn    func(ctx context.Context, id string) error
}

func (m *UserSessionMock) GetSessionByUserID(ctx context.Context, userID string) (*models.UserSession, error) {
	return m.GetSessionByUserIDFn(ctx, userID)
}

func (m *UserSessionMock) CreateSession(ctx context.Context, session *models.UserSession) error {
	return m.CreateSessionFn(ctx, session)
}

func (m *UserSessionMock) UpdateSession(ctx context.Context, session *models.UserSession) error {
	return m.UpdateSessionFn(ctx, session)
}

func (m *UserSessionMock) UpdateSessionByUserID(ctx context.Context, userID string, session *models.UserSession) error {
	return m.UpdateSessionByUserIDFn(ctx, userID, session)
}

func (m *UserSessionMock) DeleteSession(ctx context.Context, id string) error {
	return m.DeleteSessionFn(ctx, id)
}