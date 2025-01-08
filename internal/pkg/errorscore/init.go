package errorscore

import (
	errcode "github.com/chhz0/goiam/internal/pkg/errorscore/code"
)

func init() {
	// base code register
	registerCode(errcode.ErrUndefind, 500, "The error code is not defined, please check the error code")

	registerCode(errcode.ErrSuccess, 200, "Success")
	registerCode(errcode.ErrUnknow, 500, "Unknown error")
	registerCode(errcode.ErrBind, 400, "Bind error")
	registerCode(errcode.ErrValidation, 400, "Validation error")
	registerCode(errcode.ErrTokenInvalid, 401, "Token invalid")
	registerCode(errcode.ErrPageNotFound, 404, "Page not found")

	registerCode(errcode.ErrDatabase, 500, "Database error")

	registerCode(errcode.ErrEncrypt, 401, "Encrypt error")
	registerCode(errcode.ErrSignatureInvalid, 401, "Signature invalid")
	registerCode(errcode.ErrExpired, 401, "Token expired")
	registerCode(errcode.ErrInvalidAuthHeader, 401, "Invalid authoriztion header")
	registerCode(errcode.ErrMissingHeader, 401, "Missing authoriztion header")
	registerCode(errcode.ErrPasswordIncorrect, 401, "Password incorrect")
	registerCode(errcode.ErrPermissionDenied, 403, "Permission denied")

	registerCode(errcode.ErrEncodingFailed, 500, "Encoding failed")
	registerCode(errcode.ErrDecodingFailed, 500, "Decoding failed")
	registerCode(errcode.ErrInvalidJSON, 500, "Invalid JSON")
	registerCode(errcode.ErrEncodingJSON, 500, "Encoding JSON failed")
	registerCode(errcode.ErrDecodingJSON, 500, "Decoding JSON failed")
	registerCode(errcode.ErrInvalidYaml, 500, "Invalid YAML")
	registerCode(errcode.ErrEncodingYaml, 500, "Encoding YAML failed")
	registerCode(errcode.ErrDecodingYaml, 500, "Decoding YAML failed")

	// api server code register
	registerCode(errcode.ErrUserNotFound, 404, "User not found")
	registerCode(errcode.ErrUserAlreadyExist, 400, "User already exist")
	registerCode(errcode.ErrReachMaxCount, 400, "Reach max count")
	registerCode(errcode.ErrSecretNotFound, 404, "Secret not found")
	registerCode(errcode.ErrPolicyNotFound, 404, "Policy not found")
}
