package helpers

import (
	"Golabi-boilerplate/models"
	"Golabi-boilerplate/packages/jwt"

	"github.com/gofiber/fiber/v2"
)

func GetUser(ctx *fiber.Ctx) *models.User {
	return ctx.Locals(jwt.UserLocalKey).(*models.User)
}
