package log

import (
	"os"

	"go.uber.org/zap/zapcore"
)

func defaultZapOptions() []ZapOption {
	return []ZapOption{
		WithCaller(true),
		Development(),
		ErrorOutput(zapcore.AddSync(os.Stderr)),
		AddStacktrace(zapcore.PanicLevel),
	}
}
