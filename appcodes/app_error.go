package appcodes

import (
	"fmt"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`

	// Don't set `Details` directly.
	// global AppError variables(InvalidArgumentError ...) are shared across the system, so we don't want to setting Details directly
	Details map[string]any `json:"details,omitempty"`
}

func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Details: nil,
	}
}

func NewAppErrorVars(code int, message string, details map[string]any) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

// Exported method for interacting with appError
func (e *AppError) Error() string {
	return fmt.Sprintf("[%v;%v;%+v]", e.Code, e.Message, e.Details)
}

func (e *AppError) ClonedWithDetails(details map[string]any) *AppError {
	return NewAppErrorVars(e.Code, e.Message, details)
}
