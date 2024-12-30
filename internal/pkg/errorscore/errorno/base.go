package errorno

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

const (
	ErrSignatureInvalid int = iota + 100201
	ErrInvalidAuthHeader
)

const ErrDatabase int = iota + 100101