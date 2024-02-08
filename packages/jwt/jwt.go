package jwt

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type GeneratedToken struct {
	Value  string
	Expire time.Time
}

const (
	AccessTokenCookieName  string = "jwt_access_token"
	RefreshTokenCookieName string = "jwt_refresh_token"
	UserLocalKey           string = "user_local"
)

func GenerateToken(userID uint, expire int64) (string, error) {
	// Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": expire,
	})

	// Sign Token
	tokenString, signTokenErr := token.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if signTokenErr != nil {
		return "", signTokenErr
	}

	return tokenString, nil
}

func DecodeToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(viper.GetString("JWT_SECRET")), nil
	})
}

func GenerateAccessAndRefreshTokens(userID uint) (*GeneratedToken, *GeneratedToken, error) {
	accessTokenExpiration := time.Now().Add(time.Hour * 24 * 7) // 7 days
	accessToken, accessTokenErr := GenerateToken(userID, accessTokenExpiration.Unix())
	if accessTokenErr != nil {
		return nil, nil, accessTokenErr
	}

	refreshTokenExpiration := time.Now().Add(time.Hour * 24 * 30) // 30 days
	refreshToken, refreshTokenErr := GenerateToken(userID, refreshTokenExpiration.Unix())
	if refreshTokenErr != nil {
		return nil, nil, refreshTokenErr

	}

	return &GeneratedToken{
			Value:  accessToken,
			Expire: accessTokenExpiration,
		}, &GeneratedToken{
			Value:  refreshToken,
			Expire: refreshTokenExpiration,
		}, nil
}

func SetGeneratedTokensInCookie(ctx *fiber.Ctx, accessToken *GeneratedToken, refreshToken *GeneratedToken) {
	ctx.Cookie(&fiber.Cookie{
		Name:     AccessTokenCookieName,
		Value:    accessToken.Value,
		Expires:  accessToken.Expire,
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Lax",
		Domain:   "", // add your front-end domain
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     RefreshTokenCookieName,
		Value:    refreshToken.Value,
		Expires:  refreshToken.Expire,
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Lax", // add your front-end domain
	})
}
