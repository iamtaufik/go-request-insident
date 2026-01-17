package handlers

import (
	"be-request-insident/internal/usecase"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct{
	UseCase *usecase.AuthUsecase
}

func NewAuthHandler(usecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{UseCase: usecase}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	accessToken, refreshToken, err := h.UseCase.Login(c.Context(), req.Username, req.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh-token",
		Value:    refreshToken,
		HTTPOnly: true,
		Path:   "/",
		Secure: true,
		SameSite: fiber.CookieSameSiteNoneMode,
		Expires: time.Now().Add(7 * 24 * time.Hour), // 1 week
	})

	c.Cookie(&fiber.Cookie{
		Name: "access-token",
		Value: accessToken,
		HTTPOnly: true,
		Path:  "/",
		Secure: true,
		SameSite: fiber.CookieSameSiteNoneMode,
		Expires: time.Now().Add(10 * time.Minute), // 10 minutes
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "login successful",
	})
}

func (h *AuthHandler) Me(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	user, err := h.UseCase.Me(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch user data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}