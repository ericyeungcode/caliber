package appcodes

import (
	"fmt"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`

	// make details private, so that it can't be set via global AppError variables directly
	// global AppError variables(InvalidArgumentError ...) are shared across the system, so we don't want to allow setting details directly
	details map[string]any `json:"details,omitempty"`
}

func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		details: nil,
	}
}

func NewAppErrorVars(code int, message string, details map[string]any) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		details: details,
	}
}

// Exported method for interacting with appError
func (e *AppError) Error() string {
	return fmt.Sprintf("[%v;%v;%+v]", e.Code, e.Message, e.details)
}

func (e *AppError) ClonedWithDetails(details map[string]any) *AppError {
	return NewAppErrorVars(e.Code, e.Message, details)
}
