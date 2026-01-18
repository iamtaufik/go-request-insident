package handlers

import (
	"be-request-insident/internal/dto"
	"be-request-insident/internal/models"
	"be-request-insident/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type ServiceRequestHandler struct {
	UseCase *usecase.ServiceRequestUsecase	
}

func NewServiceRequestHandler(usecase *usecase.ServiceRequestUsecase) *ServiceRequestHandler {
	return &ServiceRequestHandler{UseCase: usecase}
}

func (h *ServiceRequestHandler) ListServiceRequests(c *fiber.Ctx) error {
	serviceRequests, err := h.UseCase.ListServiceRequests(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch service requests",
		})
	}
	return c.JSON(serviceRequests)
}

func (h *ServiceRequestHandler) DraftServiceRequest(c *fiber.Ctx) error {
	var req struct {
		RequestType  string `json:"request_type"`
		ReportedFrom string `json:"reported_from"`
		SLAType      string `json:"sla_type"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err := h.UseCase.DraftServiceRequest(c.Context(), &dto.DraftServiceRequestPayload{
		RequestType:  req.RequestType,
		ReportedFrom: req.ReportedFrom,
		SLAType:      req.SLAType,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to draft service request",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "service request drafted successfully",
	})
}

func (h *ServiceRequestHandler) UpdateServiceRequest(c *fiber.Ctx) error {
	serviceRequestId := c.Params("id")

	var req struct {
		Company      string `json:"company"`
		Organization string `json:"organization"`
		Department   string `json:"department"`
		Impact       string `json:"impact"`
		Urgency      string `json:"urgency"`
		Priority     string `json:"priority"`
		Summary      string `json:"summary"`
		Note         string `json:"note"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err := h.UseCase.UpdateServiceRequest(c.Context(), serviceRequestId, &dto.UpdateServiceRequestPayload{
		Company: req.Company,
		Organization: req.Organization,
		Department: req.Department,
		Impact: req.Impact,
		Urgency: req.Urgency,
		Priority: req.Priority,
		Summary: req.Summary,
		Note: req.Note,
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "service request updated successfully",
	})
}


func (h *ServiceRequestHandler) GetServiceRequestByID(c *fiber.Ctx) error {
	serviceRequestId := c.Params("id")

	serviceRequest,  err := h.UseCase.GetServiceRequestByID(c.Context(), serviceRequestId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch service request",
		})
	}

	return c.JSON(serviceRequest)
}


func (h *ServiceRequestHandler) AttachFileToServiceRequest(c *fiber.Ctx) error {
	serviceRequestId := c.Params("id")

	fileHeader, err := c.FormFile("attachment")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to get attachment file",
		})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to open attachment file",
		})
	}
	defer file.Close()

	attachment := &models.ServiceRequestAttachment{
		FileName: fileHeader.Filename,
		FileSize: fileHeader.Size,
		MimeType: fileHeader.Header.Get("Content-Type"),
		FileURL: "", // In a real application, you would upload the file to a storage service and get the URL
	}

	err = h.UseCase.AttachFileToServiceRequest(c.Context(), serviceRequestId, attachment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to attach file to service request",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "file attached to service request successfully",
	})
}


func (h *ServiceRequestHandler) GetAttachemnts(c *fiber.Ctx) error {
	serviceRequestId := c.Params("id")

	attachemnts, err := h.UseCase.GetAttachmentsByServiceRequestID(c.Context(), serviceRequestId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch attachemnts",
		})
	}

	return c.JSON(attachemnts)
}

func (h *ServiceRequestHandler) DeleteServiceRequest(c *fiber.Ctx) error {
	serviceRequestId := c.Params("id")

	err := h.UseCase.DeleteServiceRequest(c.Context(), serviceRequestId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete service request",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "service request deleted successfully",
	})
}