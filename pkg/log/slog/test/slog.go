package main

import (
	"log/slog"

	log "github.com/chhz0/goiam/pkg/log/slog"
)

type contextKey string

const contextKeyWithValue = contextKey("Context")

type User struct {
	ID       int
	Name     string
	Password string
}

func (u *User) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("id", u.ID),
		slog.String("name", u.Name),
	)
}

func main() {
	// slog example
	{
		// slog.SetLogLoggerLevel(slog.LevelInfo)
		// slog.Debug("debug message")
		// slog.Info("info message")
		// slog.Warn("warn message")
		// slog.Error("error message")

		// l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		// 	AddSource:   true,
		// 	Level:       slog.LevelInfo,
		// 	ReplaceAttr: nil,
		// }))
		// l.Debug("debug message")
		// l.Info("info message")
		// l.Warn("warn message")
		// l.Error("error message")

		// l.LogAttrs(context.Background(), slog.LevelInfo, "info message", slog.Any("key", "value"))

		// slog.SetDefault(l)
		// slog.Info("info message")
	}
	// 二次封装 SlogLogger
	{
		// log.Info("info message")

		l := log.New(log.LevelInfo)
		// l.Info("info message")
		// l.Infof("info message %s", "info")
		// l.InfoContext(context.WithValue(context.Background(), contextKeyWithValue, "withValue"), "info message")

		// l.Debug("debug message")
		// l.Debugf("debug message %s", "debuf")
		// l.DebugContext(context.WithValue(context.Background(), contextKeyWithValue, "withValue"), "debug message")

		// l.Trace("trace message")
		// l.Tracef("error message %s", "trace")
		// l.TraceContext(context.WithValue(context.Background(), contextKeyWithValue, "withValue"), "trace message")

		// l.Warn("warn message")
		// l.Warnf("warn message %s", "warn")
		// l.WarnContext(context.WithValue(context.Background(), contextKeyWithValue, "withValue"), "warn message")

		// l.Error("error message")
		// l.Errorf("error message %s", "error")
		// l.ErrorContext(context.WithValue(context.Background(), contextKeyWithValue, "withValue"), "error message")

		// l.SetLevel(log.LevelDebug)
		// l.Debug("debug message")

		// l.Info("get log level", log.Any("l.getloglevel", l.GetLogLevel()))

		user := &User{
			ID:       123,
			Name:     "chlluanma",
			Password: "pwd",
		}
		l.Info("get log level", "user", user)

		l.V(log.LevelWarn).Info("debug message")
	}
}
