package logger

import (
	"io"
	"os"

	"github.com/chhz0/goiam/pkg/log"
	"go.uber.org/zap/zapcore"
)

const (
	KeyRequestID = "requestID"
	KeyUserID    = "userID"
	KeyUsername  = "username"
)

func NewLogger() {
	logOpts := []log.ZapOption{
		log.WithCaller(true),
		log.AddCallerSkip(3),
		log.Development(),
		log.AddStacktrace(zapcore.PanicLevel),
		log.ErrorOutput(zapcore.AddSync(log.OpenLogFile("../../log/iam-api-zap-errors.log"))),
	}
	l := log.NewLogger(
		logOutput,
		log.InfoLevel,
		log.JsonEncoder,
		logOpts...,
	)

	log.ReplaceDefault(l)
}

func NewTeeLogger() {
	l := log.NewTeeLogger(
		[]log.TeeOption{},
		log.JsonEncoder,
	)

	log.ReplaceDefault(l)
}

func logOutput() io.Writer {
	return os.Stdout
}
