package slog

// var defaultLogger atomic.Pointer[SlogLogger]

// func Default() *SlogLogger {
// 	return defaultLogger.Load()
// }

// func init() {
// 	defaultLogger.Store(New(LevelDebug))
// }

// func Info(msg string, args ...any) {
// 	Default().Info(msg, args...)
// }

// // func Infof(format string, args ...any) {
// // 	std.Infof(format, args...)
// // }

// // func InfoContext(ctx context.Context, msg string, args ...any) {
// // 	std.InfoContext(ctx, msg, args...)
// // }
