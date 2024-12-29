package logger

import (
	"testing"

	"github.com/chhz0/goiam/pkg/log"
)

func TestLogger_Info(t *testing.T) {
	NewLogger()
	log.Info("log info test", log.String("key", "value"))
}
