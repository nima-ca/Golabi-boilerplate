package helpers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// HTTPError represents an error with an associated HTTP status code.
type HTTPError struct {
	Code     int
	Messages *[]string
}

type APIResponse[T any] struct {
	Status int       `json:"statusCode"`
	Data   *T        `json:"data"` // Use a pointer to T to allow nil values
	Errors *[]string `json:"errors"`
}

type PaginatedResponse struct {
	Page       int   `json:"page"`
	Count      int   `json:"count"`
	TotalCount int64 `json:"totalCount"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

func NewAPIResponse[T any](statusCode int, data *T, errors []string) APIResponse[T] {
	return APIResponse[T]{
		Status: statusCode,
		Data:   data,
		Errors: &errors,
	}
}

func OkResponse[T any](data T) APIResponse[T] {
	return NewAPIResponse[T](http.StatusOK, &data, nil)
}

func ErrorResponse[T any](statusCode int, errors []string) APIResponse[T] {
	return NewAPIResponse[T](statusCode, nil, errors)
}

func SendErrorResponse(ctx *fiber.Ctx, statusCode int, messages ...string) error {
	return ctx.Status(statusCode).JSON(ErrorResponse[any](statusCode, messages))
}

// NewHTTPError creates a new HTTPError instance.
func NewHTTPError(code int, messages ...string) *HTTPError {
	return &HTTPError{
		Code:     code,
		Messages: &messages,
	}
}

func HandleHTTPErrors(ctx *fiber.Ctx, err *HTTPError) error {
	return SendErrorResponse(ctx, err.Code, *err.Messages...)
}
