package usecase

import (
	"be-request-insident/internal/dto"
	"be-request-insident/internal/models"
	"be-request-insident/internal/repository"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type ServiceRequestUsecase struct {
	serviceRequestRepo repository.ServiceRequestRepository
	attachmentRepo repository.AttachmentRepository
}

func NewServiceRequestUsecase(serviceRequestRepo repository.ServiceRequestRepository, attachmentRepo repository.AttachmentRepository) *ServiceRequestUsecase {
	return &ServiceRequestUsecase{serviceRequestRepo: serviceRequestRepo, attachmentRepo: attachmentRepo}
}

func slaPrefix(slaType string) string {
	switch strings.ToUpper(strings.TrimSpace(slaType)) {
	case "REQUEST":
		return "REQ"
	default:
		return strings.ToUpper(slaType)[:3]
	}
}

func generateCodeBySLAType(slaType string, seq int) string {
	return fmt.Sprintf("%s-%04d", slaPrefix(slaType), seq)
}

func (u *ServiceRequestUsecase) ListServiceRequests(ctx context.Context) ([]*models.ServiceRequest, error) {
	return u.serviceRequestRepo.GetServiceRequests(ctx)
}

func (u *ServiceRequestUsecase) DraftServiceRequest(ctx context.Context, draftServiceRequest *dto.DraftServiceRequestPayload) error {
	
	
	serviceRequest := &models.ServiceRequest{
		ID: uuid.New().String(),
		RequestType: draftServiceRequest.RequestType,
		ReportedFrom: draftServiceRequest.ReportedFrom,
		SLAType: draftServiceRequest.SLAType,
		Status: "DRAFT",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	nextSeq, err := u.serviceRequestRepo.GetNextSequenceByRequestType(ctx, draftServiceRequest.RequestType)
	if err != nil {
		return err
	}

	serviceRequest.RequestCode = generateCodeBySLAType(draftServiceRequest.RequestType, nextSeq)

	return u.serviceRequestRepo.CreateServiceRequest(ctx, serviceRequest)
}

func (u *ServiceRequestUsecase) UpdateServiceRequest(ctx context.Context, serviceReuestID string,  payload *dto.UpdateServiceRequestPayload) error {
	serviceRequest, err := u.serviceRequestRepo.GetServiceRequestByID(ctx, serviceReuestID)

	if err != nil {
		return  errors.New("not found")
	}

	serviceRequest.Company = &payload.Company
	serviceRequest.Organization = &payload.Organization
	serviceRequest.Department = &payload.Department
	serviceRequest.Impact = &payload.Impact
	serviceRequest.Urgency = &payload.Urgency
	serviceRequest.Priority = &payload.Priority
	serviceRequest.Summary = &payload.Summary
	serviceRequest.Note = &payload.Note

	err = u.serviceRequestRepo.UpdateServiceRequest(ctx, serviceRequest)
	if err != nil {
		return err
	}

	return nil
}

func (u *ServiceRequestUsecase) GetServiceRequestByID(ctx context.Context, serviceRequestID string) (*models.ServiceRequest, error) {
	return u.serviceRequestRepo.GetServiceRequestByID(ctx, serviceRequestID)
}

func (u *ServiceRequestUsecase) AttachFileToServiceRequest(ctx context.Context, serviceRequestID string, attachment *models.ServiceRequestAttachment) error {
	serviceRequest, err := u.serviceRequestRepo.GetServiceRequestByID(ctx, serviceRequestID)
	if err != nil {
		return errors.New("service request not found")
	}
	attachment.ID = uuid.New().String()
	attachment.ServiceRequestID = serviceRequest.ID
	attachment.UploadedAt = time.Now()

	return u.attachmentRepo.CreateAttachment(ctx, attachment)
}

func (u *ServiceRequestUsecase) GetAttachmentsByServiceRequestID(ctx context.Context, serviceRequestID string) ([]*models.ServiceRequestAttachment, error) {
	return u.attachmentRepo.GetAttachmentsByRequestID(ctx, serviceRequestID)
}

func (u *ServiceRequestUsecase) DeleteServiceRequest(ctx context.Context, serviceRequestID string) error {
	serviceRequest, err := u.serviceRequestRepo.GetServiceRequestByID(ctx, serviceRequestID)
	if err != nil {
		return errors.New("service request not found")
	}

	now := time.Now()
	serviceRequest.DeletedAt = &now

	return u.serviceRequestRepo.UpdateServiceRequest(ctx, serviceRequest)
}