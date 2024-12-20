package ginserver

import (
	"os"
	"syscall"
)

var shutdownSignal = []os.Signal{syscall.SIGTERM, syscall.SIGINT}
