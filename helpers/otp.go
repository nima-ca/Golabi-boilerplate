package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	OTPExpireTime time.Duration = time.Minute * 2
)

func GenerateOTP() string {
	otp := rand.Intn(900000) + 100000
	return fmt.Sprint(otp)
}

func GetOTPExpireTime() time.Time {
	return time.Now().Add(OTPExpireTime)
}
