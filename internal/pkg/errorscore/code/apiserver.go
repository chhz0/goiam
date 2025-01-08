package errcode

// iam-apiserver: user errors.
const (
	ErrUserNotFound int = iota + 110001

	ErrUserAlreadyExist
)

// iam-apiserver: secret errors.
const (
	ErrReachMaxCount int = iota + 110101

	ErrSecretNotFound
)

// iam-apiserver: policy errors.
const (
	ErrPolicyNotFound int = iota + 110201
)
