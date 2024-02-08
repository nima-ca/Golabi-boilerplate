package dto

type CreateUserDto struct {
	Name     string
	Email    string
	Password string
	Otp      string
}

type RegisterHandlerBody struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

type LoginHandlerBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

type UserInfo struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	IsVerified bool   `json:"isVerified"`
}
type LoginHandlerResponseDto struct {
	User UserInfo `json:"user"`
}

type VerifyHandlerBody struct {
	Email string `json:"email" validate:"required,email"`
	Otp   string `json:"otp" validate:"required,len=6"`
}

type ResendOTPHandlerBody struct {
	Email string `json:"email" validate:"required,email"`
}
