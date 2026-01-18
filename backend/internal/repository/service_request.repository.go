package repository

import (
	"be-request-insident/internal/models"
	"context"

	"gorm.io/gorm"
)

type ServiceRequestRepository interface {
	GetServiceRequests(ctx context.Context) (*[]models.ServiceRequest, error)
	GetServiceRequestByID(ctx context.Context, id string) (*models.ServiceRequest, error)
	CreateServiceRequest(ctx context.Context, serviceRequest *models.ServiceRequest) error
	UpdateServiceRequest(ctx context.Context, serviceRequest *models.ServiceRequest) error
	GetNextSequenceByRequestType(ctx context.Context, requestType string) (int, error)
}

type serviceRequestRepository struct {
	DB *gorm.DB
	cache Cache
}

func NewServiceRequestRepository(db *gorm.DB, cache Cache) ServiceRequestRepository {
	return &serviceRequestRepository{DB: db, cache: cache}
}


func (r *serviceRequestRepository) GetServiceRequests(ctx context.Context) (*[]models.ServiceRequest, error) {
	var serviceRequests []models.ServiceRequest

	if err := r.DB.Find(&serviceRequests).Error; err != nil {
		return nil, err
	}
	return &serviceRequests, nil
}

func (r *serviceRequestRepository) CreateServiceRequest(ctx context.Context, serviceRequest *models.ServiceRequest) error {
	if err := r.DB.Create(serviceRequest).Error; err != nil {
		return err
	}
	return nil
}

func (r *serviceRequestRepository) UpdateServiceRequest(ctx context.Context, serviceRequest *models.ServiceRequest) error {
	if err := r.DB.Save(serviceRequest).Error; err != nil {
		return err
	}
	return nil
}

func (r *serviceRequestRepository) GetNextSequenceByRequestType(ctx context.Context, requestType string) (int, error) {
	var maxSeq int

	err := r.DB.WithContext(ctx).
		Model(&models.ServiceRequest{}).
		Where("request_type = ?", requestType).
		Select("COUNT(id)").
		Scan(&maxSeq).Error
	if err != nil {
		return 0, err
	}

	return maxSeq + 1, nil
}

func (r *serviceRequestRepository) GetServiceRequestByID(ctx context.Context, id string) (*models.ServiceRequest, error) {
	var requestUser models.ServiceRequest

	if err := r.DB.First(&requestUser, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &requestUser, nil
}