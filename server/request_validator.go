package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	MessageError = "request failed with validations errors"
)

type ExclusiveErrorResponse struct {
	Message string      `json:"message"`
	Field   string      `json:"field"`
	Tag     string      `json:"tag"`
	Value   interface{} `json:"value"`
}

type ErrorResponse struct {
	Errors  []ExclusiveErrorResponse `json:"errors"`
	Message string                   `json:"message"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

type RequestValidator struct {
	validator *validator.Validate
}

func NewRequestValidator() *RequestValidator {
	validate := validator.New()
	validate.RegisterValidation("notblank", validators.NotBlank)
	return &RequestValidator{validate}
}

func (cv *RequestValidator) Validate(i interface{}) error {

	errors := cv.validator.Struct(i)
	if errors != nil {
		validationErrors := errors.(validator.ValidationErrors)
		return echo.NewHTTPError(http.StatusBadRequest, createErrorResponse(validationErrors))
	}
	return nil
}

func createErrorResponse(validationErrors validator.ValidationErrors) ErrorResponse {
	var errorResponses []ExclusiveErrorResponse
	for _, err := range validationErrors {
		errorResponses = append(errorResponses, ExclusiveErrorResponse{
			Message: "Field validation failed",
			Field:   err.StructNamespace(),
			Tag:     err.ActualTag(),
			Value:   err.Value(),
		})
	}

	return ErrorResponse{
		Errors:  errorResponses,
		Message: MessageError,
	}
}
