package routes

import (
	"be-request-insident/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	AuthHandler *handlers.AuthHandler
}

func RegisterRoutes(app *fiber.App, cfg *RouteConfig) {
	api := app.Group("/api")

	RegisterAuthRoutes(api, cfg.AuthHandler)
}