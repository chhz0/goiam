package errorscore

import "github.com/chhz0/goiam/internal/pkg/constants/errorsno"

func init() {
	registerCode(errorsno.ErrUndefind, 200, "The error code is not defined, please check the error code")

	registerCode(errorsno.ErrSuccess, 200, "Success")
	registerCode(errorsno.ErrPageNotFound, 404, "Page not found")
	registerCode(errorsno.ErrSignatureInvalid, 401, "Signature invalid")
	registerCode(errorsno.ErrInvalidAuthHeader, 401, "Invalid auth header")
}
