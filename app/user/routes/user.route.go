package routes

import (
	"Golabi-boilerplate/app/user/controller"

	"github.com/gofiber/fiber/v2"
)

func Register(router *fiber.App) {
	userGroup := router.Group("/user")

	userGroup.Post("/register", controller.RegisterHandler)
	userGroup.Post("/login", controller.LoginHandler)
	userGroup.Post("/verify", controller.VerifyUserHandler)
	userGroup.Post("/resendOTP", controller.ResendOTPHandler)
}
