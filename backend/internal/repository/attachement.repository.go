package repository

import (
	"be-request-insident/internal/models"
	"context"

	"gorm.io/gorm"
)

type AttachmentRepository interface {
	GetAttachmentByID(ctx context.Context, id string) (*models.ServiceRequestAttachment, error)
	GetAttachmentsByRequestID(ctx context.Context, requestID string) ([]*models.ServiceRequestAttachment, error)
	CreateAttachment(ctx context.Context, attachment *models.ServiceRequestAttachment) error
	DeleteAttachment(ctx context.Context, id string) error
}

type attachemntRepository struct {
	DB *gorm.DB
	cache Cache
}

func NewAttachmentRepository(db *gorm.DB, cache Cache) AttachmentRepository {
	return &attachemntRepository{DB: db, cache: cache}
}

func (r *attachemntRepository) GetAttachmentByID(ctx context.Context, id string) (*models.ServiceRequestAttachment, error) {
	var attachment models.ServiceRequestAttachment

	if err := r.DB.First(&attachment, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &attachment, nil
}


func (r *attachemntRepository) CreateAttachment(ctx context.Context, attachment *models.ServiceRequestAttachment) error {
	if err := r.DB.Create(attachment).Error; err != nil {
		return err
	}
	return nil
}


func (r *attachemntRepository) DeleteAttachment(ctx context.Context, id string) error {
	if err := r.DB.Delete(&models.ServiceRequestAttachment{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *attachemntRepository) GetAttachmentsByRequestID(ctx context.Context, requestID string) ([]*models.ServiceRequestAttachment, error) {
	var attachments []*models.ServiceRequestAttachment

	if err := r.DB.Where("service_request_id = ?", requestID).Find(&attachments).Error; err != nil {
		return nil, err
	}
	return attachments, nil
}