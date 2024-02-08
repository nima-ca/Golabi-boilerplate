package controller

import (
	"Golabi-boilerplate/app/user/dto"
	"Golabi-boilerplate/app/user/services"
	"Golabi-boilerplate/helpers"
	"Golabi-boilerplate/packages/errors"
	"Golabi-boilerplate/packages/jwt"
	"Golabi-boilerplate/packages/validators"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(ctx *fiber.Ctx) error {
	body, httpErr := validators.ValidatePayload[dto.RegisterHandlerBody](ctx)
	if httpErr != nil {
		return helpers.HandleHTTPErrors(ctx, httpErr)
	}

	if _, httpErr := services.Register(body); httpErr != nil {
		return helpers.HandleHTTPErrors(ctx, httpErr)
	}

	return ctx.Status(http.StatusOK).
		JSON(helpers.OkResponse[helpers.SuccessResponse](helpers.SuccessResponse{Success: true}))
}

func LoginHandler(ctx *fiber.Ctx) error {
	body, httpErr := validators.ValidatePayload[dto.LoginHandlerBody](ctx)
	if httpErr != nil {
		return helpers.HandleHTTPErrors(ctx, httpErr)
	}

	user, httpErr := services.Login(body)

	if httpErr != nil {
		return helpers.HandleHTTPErrors(ctx, httpErr)
	}

	// Generate Tokens
	accessToken, refreshToken, generateTokenErr := jwt.GenerateAccessAndRefreshTokens(user.ID)
	if generateTokenErr != nil {
		return helpers.SendErrorResponse(ctx, http.StatusInternalServerError, errors.InternalServerErrorErrorMsg)
	}

	// Set Tokens in cookies
	jwt.SetGeneratedTokensInCookie(ctx, accessToken, refreshToken)

	return ctx.Status(http.StatusOK).
		JSON(helpers.OkResponse[dto.LoginHandlerResponseDto](dto.LoginHandlerResponseDto{
			User: dto.UserInfo{
				ID:         user.ID,
				Name:       user.Name,
				Email:      user.Email,
				IsVerified: user.IsVerified,
			},
		}))
}

func VerifyUserHandler(ctx *fiber.Ctx) error {
	body, httpErr := validators.ValidatePayload[dto.VerifyHandlerBody](ctx)
	if httpErr != nil {
		return helpers.HandleHTTPErrors(ctx, httpErr)
	}

	if httpErr := services.VerifyUser(ctx, body); httpErr != nil {
		return helpers.HandleHTTPErrors(ctx, httpErr)
	}

	return ctx.Status(http.StatusOK).
		JSON(helpers.OkResponse[helpers.SuccessResponse](helpers.SuccessResponse{Success: true}))
}

func ResendOTPHandler(ctx *fiber.Ctx) error {
	body, httpErr := validators.ValidatePayload[dto.ResendOTPHandlerBody](ctx)
	if httpErr != nil {
		return helpers.HandleHTTPErrors(ctx, httpErr)
	}

	if httpErr := services.ResendOTP(ctx, body); httpErr != nil {
		return helpers.HandleHTTPErrors(ctx, httpErr)
	}

	return ctx.Status(http.StatusOK).
		JSON(helpers.OkResponse[helpers.SuccessResponse](helpers.SuccessResponse{Success: true}))
}
