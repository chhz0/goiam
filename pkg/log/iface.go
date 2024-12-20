package log

import "context"

type InfoLogger interface {
}

type Logger interface {

	// V returns an InfoLogger value for a specific verbosity level, relative to
	// this Logger.
	V() InfoLogger

	// WithValues adds some key-value pairs of context to a logger.
	WithValue() Logger

	WithContext(ctx context.Context)
	FromContext(ctx context.Context) Logger

	L(ctx context.Context) Logger
}
