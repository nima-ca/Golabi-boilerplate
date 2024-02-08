package services

import (
	"Golabi-boilerplate/app/user/repository"
	"Golabi-boilerplate/helpers"
	"Golabi-boilerplate/models"
	"Golabi-boilerplate/packages/db"
	"Golabi-boilerplate/packages/email"
	"Golabi-boilerplate/packages/errors"
	"fmt"
	"net/http"
	"time"

	userDto "Golabi-boilerplate/app/user/dto"
	userErrors "Golabi-boilerplate/app/user/errors"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(dto *userDto.RegisterHandlerBody) (*models.User, *helpers.HTTPError) {
	user, userErr := repository.FindUserByEmail(dto.Email)
	if userErr != nil {
		return nil, helpers.NewHTTPError(http.StatusInternalServerError, errors.InternalServerErrorErrorMsg)
	}

	// check if user exists
	if user != nil {
		return nil, helpers.NewHTTPError(http.StatusBadRequest, userErrors.UserAlreadyExistErrorMsg)
	}

	otp := helpers.GenerateOTP()

	_, err := repository.CreateUser(userDto.CreateUserDto{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
		Otp:      otp,
	})

	if err != nil {
		return nil, helpers.NewHTTPError(http.StatusBadRequest, userErrors.FailedToCreateUserErrorMsg)
	}

	// DOC: send OTP and not wait for the response
	go email.SendOTP(dto.Email, otp)

	return user, nil
}

func Login(dto *userDto.LoginHandlerBody) (*models.User, *helpers.HTTPError) {
	user, userErr := repository.FindUserByEmail(dto.Email)
	if userErr != nil {
		return nil, helpers.NewHTTPError(http.StatusInternalServerError, errors.InternalServerErrorErrorMsg)
	}

	// check if user exists
	if user == nil {
		return nil, helpers.NewHTTPError(http.StatusBadRequest, userErrors.InvalidEmailOrPasswordErrorMsg)
	}

	// Check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return nil, helpers.NewHTTPError(http.StatusBadRequest, userErrors.InvalidEmailOrPasswordErrorMsg)
	}

	return user, nil
}

func VerifyUser(ctx *fiber.Ctx, dto *userDto.VerifyHandlerBody) *helpers.HTTPError {
	// Find the user
	user, err := repository.FindUserByEmail(dto.Email)
	if err != nil {
		return helpers.NewHTTPError(http.StatusInternalServerError, errors.InternalServerErrorErrorMsg)
	}

	// Verify user exists
	if user == nil {
		return helpers.NewHTTPError(http.StatusBadRequest, userErrors.IncorrectEmailErrorMsg)
	}

	// Check if user is already verified
	if user.IsVerified {
		return helpers.NewHTTPError(http.StatusBadRequest, userErrors.UserAlreadyVerifiedErrorMsg)
	}

	// Verify that an OTP is sent to user
	if user.Otp == "" {
		return helpers.NewHTTPError(http.StatusBadRequest, userErrors.IncorrectEmailErrorMsg)
	}

	// Check if OTP is expired
	if time.Now().After(user.OtpExpireTime) {
		return helpers.NewHTTPError(http.StatusBadRequest, userErrors.OTPExpiredErrorMsg)
	}

	// Verify OTP
	if err := bcrypt.CompareHashAndPassword([]byte(user.Otp), []byte(dto.Otp)); err != nil {
		return helpers.NewHTTPError(http.StatusBadRequest, userErrors.IncorrectOTPErrorMsg)
	}

	result := db.GetDB().Model(&user).Updates(map[string]interface{}{"otp": "", "is_verified": true})

	if result.Error != nil || result.RowsAffected != 1 {
		return helpers.NewHTTPError(http.StatusInternalServerError, errors.InternalServerErrorErrorMsg)
	}

	return nil
}

func ResendOTP(ctx *fiber.Ctx, dto *userDto.ResendOTPHandlerBody) *helpers.HTTPError {
	// Find the user
	user, err := repository.FindUserByEmail(dto.Email)
	if err != nil {
		return helpers.NewHTTPError(http.StatusInternalServerError, errors.InternalServerErrorErrorMsg)
	}

	// Verify user exists
	if user == nil {
		return helpers.NewHTTPError(http.StatusBadRequest, userErrors.IncorrectEmailErrorMsg)
	}

	// Check if user is already verified
	if user.IsVerified {
		return helpers.NewHTTPError(http.StatusBadRequest, userErrors.UserAlreadyVerifiedErrorMsg)
	}

	if time.Now().Before(user.OtpExpireTime) {
		diff := time.Until(user.OtpExpireTime)
		return helpers.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("You should wait %d seconds for next otp",
			int(diff.Seconds())))
	}

	otp := helpers.GenerateOTP()
	// Hash the Otp before storing it
	hashedOtp, err := bcrypt.GenerateFromPassword([]byte(otp), 10)
	if err != nil {
		return helpers.NewHTTPError(http.StatusInternalServerError, errors.InternalServerErrorErrorMsg)
	}

	result := db.GetDB().Model(&user).Updates(map[string]interface{}{"otp": hashedOtp,
		"otp_expire_time": helpers.GetOTPExpireTime()})

	if result.Error != nil || result.RowsAffected != 1 {
		return helpers.NewHTTPError(http.StatusInternalServerError, errors.InternalServerErrorErrorMsg)
	}

	// DOC: send OTP and not wait for the response
	go email.SendOTP(dto.Email, otp)

	return nil
}
