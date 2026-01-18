package routes

import (
	"be-request-insident/internal/handlers"
	"be-request-insident/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(router fiber.Router, h *handlers.AuthHandler) {
	auth := router.Group("/auth")

	auth.Post("/login", h.Login)
	auth.Post("/refresh-token", h.RefreshToken)
	auth.Delete("/logout", middlewares.JWTProtected(), h.Logout)
	auth.Get("/me",middlewares.JWTProtected(), h.Me)
}