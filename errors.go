package httperror

import (
	"fmt"
	"net/http"
)

var (
	ErrStatus = fmt.Errorf("http status") // This is the base error.

	ErrStatusBadRequest                   = fmt.Errorf("%w: %d", ErrStatus, http.StatusBadRequest)
	ErrStatusUnauthorized                 = fmt.Errorf("%w: %d", ErrStatus, http.StatusUnauthorized)
	ErrStatusPaymentRequired              = fmt.Errorf("%w: %d", ErrStatus, http.StatusPaymentRequired)
	ErrStatusForbidden                    = fmt.Errorf("%w: %d", ErrStatus, http.StatusForbidden)
	ErrStatusNotFound                     = fmt.Errorf("%w: %d", ErrStatus, http.StatusNotFound)
	ErrStatusMethodNotAllowed             = fmt.Errorf("%w: %d", ErrStatus, http.StatusMethodNotAllowed)
	ErrStatusNotAcceptable                = fmt.Errorf("%w: %d", ErrStatus, http.StatusNotAcceptable)
	ErrStatusProxyAuthRequired            = fmt.Errorf("%w: %d", ErrStatus, http.StatusProxyAuthRequired)
	ErrStatusRequestTimeout               = fmt.Errorf("%w: %d", ErrStatus, http.StatusRequestTimeout)
	ErrStatusConflict                     = fmt.Errorf("%w: %d", ErrStatus, http.StatusConflict)
	ErrStatusGone                         = fmt.Errorf("%w: %d", ErrStatus, http.StatusGone)
	ErrStatusLengthRequired               = fmt.Errorf("%w: %d", ErrStatus, http.StatusLengthRequired)
	ErrStatusPreconditionFailed           = fmt.Errorf("%w: %d", ErrStatus, http.StatusPreconditionFailed)
	ErrStatusRequestEntityTooLarge        = fmt.Errorf("%w: %d", ErrStatus, http.StatusRequestEntityTooLarge)
	ErrStatusRequestURITooLong            = fmt.Errorf("%w: %d", ErrStatus, http.StatusRequestURITooLong)
	ErrStatusUnsupportedMediaType         = fmt.Errorf("%w: %d", ErrStatus, http.StatusUnsupportedMediaType)
	ErrStatusRequestedRangeNotSatisfiable = fmt.Errorf("%w: %d", ErrStatus, http.StatusRequestedRangeNotSatisfiable)
	ErrStatusExpectationFailed            = fmt.Errorf("%w: %d", ErrStatus, http.StatusExpectationFailed)
	ErrStatusTeapot                       = fmt.Errorf("%w: %d", ErrStatus, http.StatusTeapot)
	ErrStatusMisdirectedRequest           = fmt.Errorf("%w: %d", ErrStatus, http.StatusMisdirectedRequest)
	ErrStatusUnprocessableEntity          = fmt.Errorf("%w: %d", ErrStatus, http.StatusUnprocessableEntity)
	ErrStatusLocked                       = fmt.Errorf("%w: %d", ErrStatus, http.StatusLocked)
	ErrStatusFailedDependency             = fmt.Errorf("%w: %d", ErrStatus, http.StatusFailedDependency)
	ErrStatusTooEarly                     = fmt.Errorf("%w: %d", ErrStatus, http.StatusTooEarly)
	ErrStatusUpgradeRequired              = fmt.Errorf("%w: %d", ErrStatus, http.StatusUpgradeRequired)
	ErrStatusPreconditionRequired         = fmt.Errorf("%w: %d", ErrStatus, http.StatusPreconditionRequired)
	ErrStatusTooManyRequests              = fmt.Errorf("%w: %d", ErrStatus, http.StatusTooManyRequests)
	ErrStatusRequestHeaderFieldsTooLarge  = fmt.Errorf("%w: %d", ErrStatus, http.StatusRequestHeaderFieldsTooLarge)
	ErrStatusUnavailableForLegalReasons   = fmt.Errorf("%w: %d", ErrStatus, http.StatusUnavailableForLegalReasons)

	ErrStatusInternalServerError           = fmt.Errorf("%w: %d", ErrStatus, http.StatusInternalServerError)
	ErrStatusNotImplemented                = fmt.Errorf("%w: %d", ErrStatus, http.StatusNotImplemented)
	ErrStatusBadGateway                    = fmt.Errorf("%w: %d", ErrStatus, http.StatusBadGateway)
	ErrStatusServiceUnavailable            = fmt.Errorf("%w: %d", ErrStatus, http.StatusServiceUnavailable)
	ErrStatusGatewayTimeout                = fmt.Errorf("%w: %d", ErrStatus, http.StatusGatewayTimeout)
	ErrStatusHTTPVersionNotSupported       = fmt.Errorf("%w: %d", ErrStatus, http.StatusHTTPVersionNotSupported)
	ErrStatusVariantAlsoNegotiates         = fmt.Errorf("%w: %d", ErrStatus, http.StatusVariantAlsoNegotiates)
	ErrStatusInsufficientStorage           = fmt.Errorf("%w: %d", ErrStatus, http.StatusInsufficientStorage)
	ErrStatusLoopDetected                  = fmt.Errorf("%w: %d", ErrStatus, http.StatusLoopDetected)
	ErrStatusNotExtended                   = fmt.Errorf("%w: %d", ErrStatus, http.StatusNotExtended)
	ErrStatusNetworkAuthenticationRequired = fmt.Errorf("%w: %d", ErrStatus, http.StatusNetworkAuthenticationRequired)
)

var ErrorStatusMap = map[int]error{
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
	if err := ErrorStatusMap[status]; err != nil {
		return err
	}
	if status < 400 {
		return nil
	}
	return ErrStatus
}
