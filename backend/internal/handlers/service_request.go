package handlers

import (
	"be-request-insident/internal/dto"
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