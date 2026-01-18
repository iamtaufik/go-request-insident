package handlers

import (
	"be-request-insident/internal/config"
	"be-request-insident/internal/handlers/response"
	"be-request-insident/internal/usecase"
	"be-request-insident/utility"
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

	user, err := h.UseCase.Me(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch user data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": response.UserMeResponse{
			ID:        user.ID,
			Username:  user.Username,
			IsActive:  user.IsActive,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	
	err := h.UseCase.Logout(c.Context(), userID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.ClearCookie("access-token")
	c.ClearCookie("refresh-token")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "logout successful",
	}) 
}

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh-token", "")
	claims, err := utility.ParseToken(refreshToken, config.GetEnvVariable("JWT_REFRESH_KEY"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid refresh token",
		})
	}
	userID, ok := utility.GetStringClaim(claims, "user_id")
    if !ok || userID == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "invalid refresh token claims",
        })
    }


	newAccessToken, newRefreshToken, err := h.UseCase.RefreshToken(c.Context(), userID, refreshToken)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh-token",
		Value:    newRefreshToken,
		HTTPOnly: true,
		Path:   "/",
		Secure: true,
		SameSite: fiber.CookieSameSiteNoneMode,
		Expires: time.Now().Add(7 * 24 * time.Hour), // 1 week
	})

	c.Cookie(&fiber.Cookie{
		Name: "access-token",
		Value: newAccessToken,
		HTTPOnly: true,
		Path:  "/",
		Secure: true,
		SameSite: fiber.CookieSameSiteNoneMode,
		Expires: time.Now().Add(10 * time.Minute), // 10 minutes
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "token refreshed successfully",
	})
}