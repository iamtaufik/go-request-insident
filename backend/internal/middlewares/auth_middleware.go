package middlewares

import (
	"be-request-insident/utility"
	"os"

	"github.com/gofiber/fiber/v2"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := c.Cookies("access-token", "")

		if accessToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or malformed token"})
		}

		claims, err := utility.ParseToken(accessToken, os.Getenv("JWT_SECRET_KEY"))

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
		}

		userID := claims["user_id"]
		if userID == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user_id in token"})
		}

		c.Locals("user_id", userID)

		return c.Next()
	}
}