package errcode

const (
	ErrUndefind int = -999
)

const (
	// ErrSuccess - 200: OK
	ErrSuccess int = iota + 100001
	ErrUnknow
	ErrBind
	ErrValidation
	ErrTokenInvalid
	ErrPageNotFound
)

const ErrDatabase int = iota + 100101

const (
	ErrEncrypt int = iota + 100201
	ErrSignatureInvalid
	ErrExpired
	ErrInvalidAuthHeader
	ErrMissingHeader
	ErrPasswordIncorrect
	ErrPermissionDenied
)

const (
	ErrEncodingFailed int = iota + 100301
	ErrDecodingFailed
	ErrInvalidJSON
	ErrEncodingJSON
	ErrDecodingJSON
	ErrInvalidYaml
	ErrEncodingYaml
	ErrDecodingYaml
)
