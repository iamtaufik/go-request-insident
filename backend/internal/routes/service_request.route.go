package routes

import (
	"be-request-insident/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterServiceRequestRoutes(router fiber.Router, h *handlers.ServiceRequestHandler) {
	serviceRequest := router.Group("/service-request")

	serviceRequest.Post("/", h.DraftServiceRequest)
	serviceRequest.Put("/:id", h.UpdateServiceRequest)
}