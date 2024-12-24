package log

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LevelEnablerFunc func(lvl Level) bool

type TeeOption struct {
	out string
	LevelEnablerFunc
}

func newTee(tees []TeeOption, opts ...ZapOption) *zapLogger {
	cores := make([]zapcore.Core, 0, len(tees))

	for _, tee := range tees {
		var out io.Writer

		out = os.Stdout
		if tee.out != "" {
			out = openLogFile(tee.out)
		}

		encoderConfig := &zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(*encoderConfig),
			zapcore.AddSync(out),
			zap.LevelEnablerFunc(tee.LevelEnablerFunc),
		)
		cores = append(cores, core)
	}

	return &zapLogger{
		l:  zap.New(zapcore.NewTee(cores...), opts...),
		al: nil,
	}
}

func openLogFile(file string) io.Writer {
	logf, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("open log file failed")
	}
	return logf
}
