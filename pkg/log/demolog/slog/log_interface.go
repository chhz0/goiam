package slog

import "context"

type InfoLogger interface {
	Info(msg string, args ...any)
	Infof(format string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)

	Enabled() bool
}

type Logger interface {
	InfoLogger
	Debug(msg string, args ...any)
	Debugf(format string, args ...any)
	DebugContext(ctx context.Context, msg string, args ...any)

	Trace(msg string, args ...any)
	Tracef(format string, args ...any)
	TraceContext(ctx context.Context, msg string, args ...any)

	Warn(msg string, args ...any)
	Warnf(format string, args ...any)
	Warnw(ctx context.Context, msg string, args ...any)

	Error(msg string, args ...any)
	Errorf(format string, args ...any)
	Errorw(ctx context.Context, msg string, args ...any)

	// V 返回一个具有指定日志级别的信息记录器。
	// 这个方法允许动态调整日志记录的详细程度，以根据需要查看更具体或更高级别的日志信息。
	// 返回的 InfoLogger 是一个具有指定日志级别并可用于记录信息的接口。
	V(level Level) InfoLogger

	With(args ...any) Logger
	WithName(name string) Logger
}
