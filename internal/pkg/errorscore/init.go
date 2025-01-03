package errorscore

import "github.com/chhz0/goiam/internal/pkg/errorscore/errorno"

func init() {
	registerCode(errorno.ErrUndefind, 200, "The error code is not defined, please check the error code")

	registerCode(errorno.ErrSuccess, 200, "Success")
	registerCode(errorno.ErrPageNotFound, 404, "Page not found")
	registerCode(errorno.ErrSignatureInvalid, 401, "Signature invalid")
	registerCode(errorno.ErrInvalidAuthHeader, 401, "Invalid auth header")
}
