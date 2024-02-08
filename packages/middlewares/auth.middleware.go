package middlewares

import (
	"Golabi-boilerplate/app/user/repository"
	"Golabi-boilerplate/helpers"
	jwtModule "Golabi-boilerplate/packages/jwt"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const (
	InvalidToken string = "Invalid token"
	ExpiredToken string = "Token Expired"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	// Get token from request cookie
	accessTokenString := ctx.Cookies(jwtModule.AccessTokenCookieName, "")
	refreshTokenString := ctx.Cookies(jwtModule.RefreshTokenCookieName, "")

	// Make sure that Access and Refresh tokens exist
	if accessTokenString == "" || refreshTokenString == "" {
		return UnauthorizedAccess(ctx, InvalidToken)
	}

	accessToken, err := jwtModule.DecodeToken(accessTokenString)

	// DOC: we check the token expiration ahead
	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		return UnauthorizedAccess(ctx, InvalidToken)
	}

	claims, ok := accessToken.Claims.(jwt.MapClaims)
	if !ok {
		return UnauthorizedAccess(ctx, InvalidToken)
	}

	// Find User
	userID := claims["sub"].(float64)
	user, userErr := repository.FindUserById(fmt.Sprint(userID))

	// check if user exists
	if userErr != nil || user == nil {
		return UnauthorizedAccess(ctx, InvalidToken)
	}

	// DOC: check if token is expired, then check if the refresh token is valid
	// then if the refresh token was ok, it regenerates access and refresh token
	// and set them in cookies.
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		// Validate Refresh Token
		refreshToken, refErr := jwtModule.DecodeToken(refreshTokenString)
		if refErr != nil {
			return UnauthorizedAccess(ctx, InvalidToken)
		}

		// Extract Refresh Token's claims
		refreshTokenClaims, ok := refreshToken.Claims.(jwt.MapClaims)
		if !ok {
			return UnauthorizedAccess(ctx, InvalidToken)
		}

		// check if Refresh Token is expired
		if float64(time.Now().Unix()) > refreshTokenClaims["exp"].(float64) {
			return UnauthorizedAccess(ctx, ExpiredToken)
		}

		// generate new tokens
		newAccessToken, newRefreshToken, tokenGenerationErr :=
			jwtModule.GenerateAccessAndRefreshTokens(user.ID)

		if tokenGenerationErr != nil {
			return UnauthorizedAccess(ctx, InvalidToken)
		}

		jwtModule.SetGeneratedTokensInCookie(ctx, newAccessToken, newRefreshToken)
	}

	ctx.Locals(jwtModule.UserLocalKey, user)

	return ctx.Next()
}

func UnauthorizedAccess(ctx *fiber.Ctx, message string) error {
	return ctx.Status(http.StatusUnauthorized).
		JSON(helpers.ErrorResponse[any](http.StatusBadRequest, []string{
			message,
		}))
}
