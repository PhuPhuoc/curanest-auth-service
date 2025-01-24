package common

import (
	"net/http"
	"runtime"
)

type AppError struct {
	// Code: HTTP status code
	Code int `json:"code"`
	// Message: The message to be returned to the client
	Message string `json:"message"`
	// ReasonField: A message describing the error during processing
	// Example: Unable to add a new record to table ABC - Unable to update the record
	ReasonField string `json:"reason_field,omitempty"`
	// Details: A map to store detailed information for better error clarity
	// Details map[string]interface{} `json:"details,omitempty"`
	// File: The file where the error occurred
	File string `json:"file,omitempty"`
	// Line: The line in the file where the error occurred
	Line int `json:"line,omitempty"`
	// Inner: Root cause of the error - useful for debugging
	Inner string `json:"inner,omitempty"`
}

func NewAppError(code int, message string, flag_location bool) *AppError {
	if flag_location {
		file, line := getCallerInfo()
		return &AppError{
			Code:    code,
			Message: message,
			File:    file,
			Line:    line,
			// Details: make(map[string]interface{}),
		}
	} else {
		return &AppError{
			Code:    code,
			Message: message,
			// Details: make(map[string]interface{}),
		}
	}
}

func NewBadRequestError() *AppError {
	file, line := getCallerInfo()
	return &AppError{
		Message: "Bad Request",
		Code:    http.StatusBadRequest,
		File:    file,
		Line:    line,
		// Details:     make(map[string]interface{}),
	}
}

func NewInternalServerError() *AppError {
	file, line := getCallerInfo()
	return &AppError{
		Message: "Internal Server Error",
		Code:    http.StatusInternalServerError,
		File:    file,
		Line:    line,
		// Details:     make(map[string]interface{}),
	}
}

func NewUnauthorizedError() *AppError {
	return &AppError{
		Message: "Unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewForbiddenError() *AppError {
	return &AppError{
		Message: "Forbiden",
		Code:    http.StatusForbidden,
	}
}

func (e *AppError) StatusCode() int {
	return e.Code
}

func (e *AppError) WithMessage(message string) *AppError {
	e.Message = message
	return e
}

func (e *AppError) WithReason(reason string) *AppError {
	e.ReasonField = reason
	return e
}

func (e *AppError) WithInner(inner string) *AppError {
	e.Inner = inner
	return e
}

func getCallerInfo() (string, int) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown", 0
	}
	return file, line
}

func (e *AppError) Error() string {
	return e.Message + " - " + e.ReasonField + " - " + e.Inner
}

//	func (e *AppError) WithDetail(key string, value interface{}) *AppError {
//		e.Details[key] = value
//		return e
//	}
