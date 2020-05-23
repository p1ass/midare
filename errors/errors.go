package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

// ServiceError is a error struct with error code
type ServiceError struct {
	Code    ErrCode
	message string
}

// ErrCode is a internal error code type
type ErrCode int

const (
	// Unknown uses when error cause is unknown
	Unknown ErrCode = iota
	// BadRequest uses when request is incorrect
	BadRequest
	// NotFound uses when a entity is not found
	NotFound
	// Unauthorized uses when a user not authorized
	Unauthorized
	// Forbidden uses when operation is not permitted
	Forbidden
	// InvalidArgument uses when argument is invalid
	InvalidArgument
)

func (e *ServiceError) Error() string {
	return e.message
}

// New returns service error
func New(code ErrCode, format string, args ...interface{}) *ServiceError {
	return &ServiceError{Code: code, message: fmt.Sprintf(format, args...)}
}

// Wrap wraps error with stack
func Wrap(err error, format string, args ...interface{}) error {
	if len(args) == 0 {
		return errors.Wrap(err, format)
	}
	return errors.Wrap(err, fmt.Sprintf(format, args...))
}

// Cause returns the root of error
func Cause(err error) error {
	return errors.Cause(err)
}

// NewNotFound returns a not found service error
func NewNotFound(format string, args ...interface{}) *ServiceError {
	return New(NotFound, format, args...)
}

// NewBadRequest returns a bad request service error
func NewBadRequest(format string, args ...interface{}) *ServiceError {
	return New(BadRequest, format, args...)
}

// NewForbidden returns a forbidden service error
func NewForbidden(format string, args ...interface{}) *ServiceError {
	return New(Forbidden, format, args...)
}

// NewUnknown returns a unknown error
func NewUnknown(format string, args ...interface{}) *ServiceError {
	return New(Unknown, format, args...)
}

// NewInvalidArgument returns a invalid argument error
func NewInvalidArgument(format string, args ...interface{}) *ServiceError {
	return New(InvalidArgument, format, args...)
}
