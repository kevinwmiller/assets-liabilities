package errors

import (
	"fmt"
)

// ErrorWithCode is a cutom error type that contains an http status code and a supporting message
type ErrorWithCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error prints the error message as a string
func (e ErrorWithCode) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

// NewErrorWithCode creates a new ErrorWithCode
func NewErrorWithCode(code int, message string) *ErrorWithCode {
	return &ErrorWithCode{
		Code:    code,
		Message: message,
	}
}

// Error converts an error into an ErrorWithCode. If the given error is not an ErrorWithCode, a 500 code is returned with the error message
func Error(err error) *ErrorWithCode {
	rdsErr, ok := err.(*ErrorWithCode)
	if !ok {
		rdsErr = NewErrorWithCode(500, err.Error())
	}
	return rdsErr
}
