package repository

import (
	"Golabi-boilerplate/app/user/dto"
	"Golabi-boilerplate/helpers"
	"Golabi-boilerplate/models"
	"Golabi-boilerplate/packages/db"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(createUserDto dto.CreateUserDto) (*models.User, error) {
	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUserDto.Password), 10)
	if err != nil {
		return nil, err
	}

	// Hash the Otp before storing it
	hashedOtp, err := bcrypt.GenerateFromPassword([]byte(createUserDto.Otp), 10)
	if err != nil {
		return nil, err
	}

	// create user
	user := models.User{
		Name:          createUserDto.Name,
		Email:         createUserDto.Email,
		Password:      string(hashedPassword),
		IsVerified:    false,
		Otp:           string(hashedOtp),
		OtpExpireTime: helpers.GetOTPExpireTime(),
	}

	// save user in DB
	result := db.GetDB().Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func FindUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	err := db.GetDB().Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func FindUserById(id string) (*models.User, error) {
	user := models.User{}
	err := db.GetDB().Where("id = ?", id).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
