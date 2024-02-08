package validators

import (
	"Golabi-boilerplate/helpers"
	"Golabi-boilerplate/packages/errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidatePayload[T any](ctx *fiber.Ctx) (*T, *helpers.HTTPError) {
	// Get Body of the Request
	var body T
	if err := ctx.BodyParser(&body); err != nil {
		return nil,
			helpers.NewHTTPError(http.StatusBadRequest, errors.FailedToParseBodyErrorMsg)
	}

	// Validate Body of request
	if errs := Validate(body); len(errs) > 0 {
		fmt.Println("hiii")
		return nil,
			helpers.NewHTTPError(http.StatusBadRequest, GetValidationErrors(errs)...)
	}

	return &body, nil
}

func ValidateSlice(slice interface{}) *helpers.HTTPError {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return helpers.NewHTTPError(
			http.StatusBadRequest,
			"Input is not slice",
		)

	}

	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i).Interface()
		if errs := Validate(item); len(errs) > 0 {
			return helpers.NewHTTPError(
				http.StatusBadRequest,
				GetValidationErrors(errs)...,
			)
		}
	}

	return nil
}

func Validate(data interface{}) []ValidationErrorResponse {
	validationErrors := []ValidationErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ValidationErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Param = err.Param()       // Export field parameter

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func GetValidationErrors(errors []ValidationErrorResponse) []string {

	var errorsSlice []string

	for _, err := range errors {
		// Customize error messages based on the field and tag
		switch err.Tag {
		case "required":
			errorsSlice = append(errorsSlice, fmt.Sprintf("%s is required.",
				err.FailedField))
		case "email":
			errorsSlice = append(errorsSlice, fmt.Sprintf("%s must be a valid email address.",
				err.FailedField))
		case "gte":
			errorsSlice = append(errorsSlice, fmt.Sprintf("%s must be greater or equal to %s.",
				err.FailedField, err.Param))
		case "lte":
			errorsSlice = append(errorsSlice, fmt.Sprintf("%s must be less than or equal to %s.",
				err.FailedField, err.Param))
		case "min":
			errorsSlice = append(errorsSlice, fmt.Sprintf("%s must be at least %s.",
				err.FailedField, err.Param))
		case "max":
			errorsSlice = append(errorsSlice, fmt.Sprintf("%s must be at most %s.",
				err.FailedField, err.Param))
		case "len":
			errorsSlice = append(errorsSlice, fmt.Sprintf("%s must have exactly %s characters.",
				err.FailedField, err.Param))
		case "dateformat":
			errorsSlice = append(errorsSlice, fmt.Sprintf("%s must be in a valid format.", err.FailedField))
		default:
			// Use the default error message
			errorsSlice = append(errorsSlice, fmt.Sprintf("Validation failed for field %s with tag %s.",
				err.FailedField, err.Tag))
		}
	}

	return errorsSlice
}

func RegisterValidators() {
	validate.RegisterValidation("dateformat", DateFormatValidation)
}

func DateFormatValidation(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	layout := fl.Param()

	_, err := time.Parse(layout, dateStr)
	return err == nil
}
