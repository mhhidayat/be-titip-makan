package jsonutil

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type response[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func ErrorResponse(message string) response[string] {
	return response[string]{
		Status:  false,
		Message: message,
		Data:    "",
	}
}

func SuccessResponse[T any](message string, data T) response[T] {
	return response[T]{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func MappingErrors(validationErrors validator.ValidationErrors) map[string]string {
	errors := make(map[string]string)
	for _, fieldErr := range validationErrors {
		fieldName := strings.ToLower(fieldErr.Field())
		tag := fieldErr.Tag()
		param := fieldErr.Param()

		var message string

		// Buat pesan user-friendly berdasarkan tag
		switch tag {
		case "required":
			message = fmt.Sprintf("Field %s is required", fieldName)
		case "min":
			message = fmt.Sprintf("Field %s must be at least %s characters", fieldName, param)
		case "max":
			message = fmt.Sprintf("Field %s must be at most %s characters", fieldName, param)
		case "email":
			message = fmt.Sprintf("Field %s must be a valid email", fieldName)
		case "gte":
			message = fmt.Sprintf("Field %s must be greater than or equal to %s", fieldName, param)
		case "lte":
			message = fmt.Sprintf("Field %s must be less than or equal to %s", fieldName, param)
		default:
			message = fmt.Sprintf("Field %s failed on the %s rule", fieldName, tag)
		}

		errors[fieldName] = message
	}
	return errors
}

func ValidationErrorResponse[T any](data T) response[T] {
	return response[T]{
		Status:  false,
		Message: "Validation failed",
		Data:    data,
	}
}
