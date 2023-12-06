package httperror

import (
	"errors"
	"fmt"
	"net/http"
)

// Error is an HTTP error.
type Error struct {
	code int // This is the HTTP status code.
}

// Error returns a string representation of the error
func (e *Error) Error() string {
	return fmt.Sprintf("http status: %d (%s)", e.code, http.StatusText(e.code))
}

// Code returns the HTTP status code.
func (e *Error) Code() int {
	return e.code
}

// Unwrap returns the base error type, which is ErrStatus
func (e *Error) Unwrap() error {
	return ErrStatus
}

// newError returns a new error for the given status code.
// We will use this internally to generate all of the ErrStatus* variables.
func newError(code int) error {
	return &Error{
		code: code,
	}
}

var (
	ErrStatus = fmt.Errorf("http status") // This is the base error.

	ErrStatusBadRequest                   = newError(http.StatusBadRequest)
	ErrStatusUnauthorized                 = newError(http.StatusUnauthorized)
	ErrStatusPaymentRequired              = newError(http.StatusPaymentRequired)
	ErrStatusForbidden                    = newError(http.StatusForbidden)
	ErrStatusNotFound                     = newError(http.StatusNotFound)
	ErrStatusMethodNotAllowed             = newError(http.StatusMethodNotAllowed)
	ErrStatusNotAcceptable                = newError(http.StatusNotAcceptable)
	ErrStatusProxyAuthRequired            = newError(http.StatusProxyAuthRequired)
	ErrStatusRequestTimeout               = newError(http.StatusRequestTimeout)
	ErrStatusConflict                     = newError(http.StatusConflict)
	ErrStatusGone                         = newError(http.StatusGone)
	ErrStatusLengthRequired               = newError(http.StatusLengthRequired)
	ErrStatusPreconditionFailed           = newError(http.StatusPreconditionFailed)
	ErrStatusRequestEntityTooLarge        = newError(http.StatusRequestEntityTooLarge)
	ErrStatusRequestURITooLong            = newError(http.StatusRequestURITooLong)
	ErrStatusUnsupportedMediaType         = newError(http.StatusUnsupportedMediaType)
	ErrStatusRequestedRangeNotSatisfiable = newError(http.StatusRequestedRangeNotSatisfiable)
	ErrStatusExpectationFailed            = newError(http.StatusExpectationFailed)
	ErrStatusTeapot                       = newError(http.StatusTeapot)
	ErrStatusMisdirectedRequest           = newError(http.StatusMisdirectedRequest)
	ErrStatusUnprocessableEntity          = newError(http.StatusUnprocessableEntity)
	ErrStatusLocked                       = newError(http.StatusLocked)
	ErrStatusFailedDependency             = newError(http.StatusFailedDependency)
	ErrStatusTooEarly                     = newError(http.StatusTooEarly)
	ErrStatusUpgradeRequired              = newError(http.StatusUpgradeRequired)
	ErrStatusPreconditionRequired         = newError(http.StatusPreconditionRequired)
	ErrStatusTooManyRequests              = newError(http.StatusTooManyRequests)
	ErrStatusRequestHeaderFieldsTooLarge  = newError(http.StatusRequestHeaderFieldsTooLarge)
	ErrStatusUnavailableForLegalReasons   = newError(http.StatusUnavailableForLegalReasons)

	ErrStatusInternalServerError           = newError(http.StatusInternalServerError)
	ErrStatusNotImplemented                = newError(http.StatusNotImplemented)
	ErrStatusBadGateway                    = newError(http.StatusBadGateway)
	ErrStatusServiceUnavailable            = newError(http.StatusServiceUnavailable)
	ErrStatusGatewayTimeout                = newError(http.StatusGatewayTimeout)
	ErrStatusHTTPVersionNotSupported       = newError(http.StatusHTTPVersionNotSupported)
	ErrStatusVariantAlsoNegotiates         = newError(http.StatusVariantAlsoNegotiates)
	ErrStatusInsufficientStorage           = newError(http.StatusInsufficientStorage)
	ErrStatusLoopDetected                  = newError(http.StatusLoopDetected)
	ErrStatusNotExtended                   = newError(http.StatusNotExtended)
	ErrStatusNetworkAuthenticationRequired = newError(http.StatusNetworkAuthenticationRequired)
)

// errorStatusMap maps an HTTP status to its corresponding ErrStatus value.
var errorStatusMap = map[int]error{
	http.StatusBadRequest:                   ErrStatusBadRequest,
	http.StatusUnauthorized:                 ErrStatusUnauthorized,
	http.StatusPaymentRequired:              ErrStatusPaymentRequired,
	http.StatusForbidden:                    ErrStatusForbidden,
	http.StatusNotFound:                     ErrStatusNotFound,
	http.StatusMethodNotAllowed:             ErrStatusMethodNotAllowed,
	http.StatusNotAcceptable:                ErrStatusNotAcceptable,
	http.StatusProxyAuthRequired:            ErrStatusProxyAuthRequired,
	http.StatusRequestTimeout:               ErrStatusRequestTimeout,
	http.StatusConflict:                     ErrStatusConflict,
	http.StatusGone:                         ErrStatusGone,
	http.StatusLengthRequired:               ErrStatusLengthRequired,
	http.StatusPreconditionFailed:           ErrStatusPreconditionFailed,
	http.StatusRequestEntityTooLarge:        ErrStatusRequestEntityTooLarge,
	http.StatusRequestURITooLong:            ErrStatusRequestURITooLong,
	http.StatusUnsupportedMediaType:         ErrStatusUnsupportedMediaType,
	http.StatusRequestedRangeNotSatisfiable: ErrStatusRequestedRangeNotSatisfiable,
	http.StatusExpectationFailed:            ErrStatusExpectationFailed,
	http.StatusTeapot:                       ErrStatusTeapot,
	http.StatusMisdirectedRequest:           ErrStatusMisdirectedRequest,
	http.StatusUnprocessableEntity:          ErrStatusUnprocessableEntity,
	http.StatusLocked:                       ErrStatusLocked,
	http.StatusFailedDependency:             ErrStatusFailedDependency,
	http.StatusTooEarly:                     ErrStatusTooEarly,
	http.StatusUpgradeRequired:              ErrStatusUpgradeRequired,
	http.StatusPreconditionRequired:         ErrStatusPreconditionRequired,
	http.StatusTooManyRequests:              ErrStatusTooManyRequests,
	http.StatusRequestHeaderFieldsTooLarge:  ErrStatusRequestHeaderFieldsTooLarge,
	http.StatusUnavailableForLegalReasons:   ErrStatusUnavailableForLegalReasons,

	http.StatusInternalServerError:           ErrStatusInternalServerError,
	http.StatusNotImplemented:                ErrStatusNotImplemented,
	http.StatusBadGateway:                    ErrStatusBadGateway,
	http.StatusServiceUnavailable:            ErrStatusServiceUnavailable,
	http.StatusGatewayTimeout:                ErrStatusGatewayTimeout,
	http.StatusHTTPVersionNotSupported:       ErrStatusHTTPVersionNotSupported,
	http.StatusVariantAlsoNegotiates:         ErrStatusVariantAlsoNegotiates,
	http.StatusInsufficientStorage:           ErrStatusInsufficientStorage,
	http.StatusLoopDetected:                  ErrStatusLoopDetected,
	http.StatusNotExtended:                   ErrStatusNotExtended,
	http.StatusNetworkAuthenticationRequired: ErrStatusNetworkAuthenticationRequired,
}

// ErrorFromStatus returns the error for the corresponding HTTP status code.
//
// If no such status could be found, ErrStatus is returned.
func ErrorFromStatus(status int) error {
	// If we have a mapped error, then return that.
	if err := errorStatusMap[status]; err != nil {
		return err
	}
	// Anything less than 400-level is not an error.
	if status < 400 {
		return nil
	}
	// Return a new error; we don't support this status code yet.
	return newError(status)
}

// StatusFromError returns the status code for the error.
//
// If the error is not an HTTP error, then this returns 0.
func StatusFromError(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Code()
	}
	return 0
}
