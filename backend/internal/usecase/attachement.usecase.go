package usecase

import (
	"be-request-insident/internal/models"
	"be-request-insident/internal/repository"
	"context"
)

type AttachmentUsecase struct {
	attachmentRepo repository.AttachmentRepository
}


func NewAttachmentUsecase(attachmentRepo repository.AttachmentRepository) *AttachmentUsecase {
	return &AttachmentUsecase{attachmentRepo: attachmentRepo}
}

func (a *AttachmentUsecase) GetAttachmentByID(ctx context.Context, id string) (*models.ServiceRequestAttachment, error) {
	return a.attachmentRepo.GetAttachmentByID(ctx, id)
}

func (a *AttachmentUsecase) GetAttachmentsByRequestID(ctx context.Context, requestID string) ([]*models.ServiceRequestAttachment, error) {
	return a.attachmentRepo.GetAttachmentsByRequestID(ctx, requestID)
}