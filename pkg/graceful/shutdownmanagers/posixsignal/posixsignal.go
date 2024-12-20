package posixsignal

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/chhz0/goiam/pkg/graceful"
)

const name = "PosixSignalManager"

type PosixSignalManager struct {
	signals []os.Signal
}

// GetName implements graceful.ShutdownManager.
func (p *PosixSignalManager) GetName() string {
	return name
}

// ShutdownFinish implements graceful.ShutdownManager.
func (p *PosixSignalManager) ShutdownFinish() error {
	return nil
}

// ShutdownStart implements graceful.ShutdownManager.
func (p *PosixSignalManager) ShutdownStart() error {
	return nil
}

// Start implements graceful.ShutdownManager.
func (p *PosixSignalManager) Start(gs graceful.GracefulShutdownIface) error {
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, p.signals...)

		<-quit

		gs.Shutdown(p)
	}()

	return nil
}

// NewPosixSignalManager creates a new PosixSignalManager
// with the given signals. If no signals are given,
// SIGINT and SIGTERM are used.
func NewPosixSignalManager(sig ...os.Signal) *PosixSignalManager {
	if len(sig) == 0 {
		sig = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	}

	return &PosixSignalManager{
		signals: sig,
	}
}

var _ graceful.ShutdownManager = (*PosixSignalManager)(nil)
