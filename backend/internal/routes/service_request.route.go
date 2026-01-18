package routes

import (
	"be-request-insident/internal/handlers"
	"be-request-insident/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterServiceRequestRoutes(router fiber.Router, h *handlers.ServiceRequestHandler) {
	serviceRequest := router.Group("/service-request")

	serviceRequest.Post("/", middlewares.JWTProtected(), h.DraftServiceRequest)
	serviceRequest.Get("/", middlewares.JWTProtected(), h.ListServiceRequests)
	serviceRequest.Get("/:id", middlewares.JWTProtected(), h.GetServiceRequestByID)
	serviceRequest.Put("/:id",middlewares.JWTProtected(), h.UpdateServiceRequest)
	serviceRequest.Delete("/:id",middlewares.JWTProtected(), h.DeleteServiceRequest)
	serviceRequest.Post("/:id/attachments", middlewares.JWTProtected(), h.AttachFileToServiceRequest)
	serviceRequest.Get("/:id/attachments", middlewares.JWTProtected(), h.GetAttachemnts)
}