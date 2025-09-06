package appcodes

import (
	"fmt"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`

	// make vars private, so that it can't be set via global AppError variables directly
	// global AppError variables(InvalidArgumentError ...) are shared across the system, so we don't want to allow setting vars directly
	vars map[string]any `json:"vars,omitempty"`
}

func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		vars:    nil,
	}
}

func NewAppErrorVars(code int, message string, ctxVars map[string]any) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		vars:    ctxVars,
	}
}

// Exported method for interacting with appError
func (e *AppError) Error() string {
	return fmt.Sprintf("[%v;%v;%+v]", e.Code, e.Message, e.vars)
}

func (e *AppError) WithCtxVars(ctxVars map[string]any) *AppError {
	v := *e // Create a copy of the current error
	v.vars = ctxVars
	return &v
}
