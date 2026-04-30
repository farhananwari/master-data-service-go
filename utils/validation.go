package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// FormatValidationError ubah error dari Gin jadi lebih clean
func FormatValidationError(err error) []string {
	var errors []string

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			field := toSnakeCase(fe.Field())
			errors = append(errors, field+" is "+fe.Tag())
		}
		return errors
	}

	// fallback kalau bukan validation error
	errors = append(errors, err.Error())
	return errors
}

// helper convert CamelCase → snake_case
func toSnakeCase(str string) string {
	var result string
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result += "_"
		}
		result += string(r)
	}
	return strings.ToLower(result)
}
