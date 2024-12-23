package graceful

import "sync"

// ShutdownCallback is a callback interface for graceful shutdown.
type ShutdownCallback interface {
	OnShutdown(string) error
}

type OnShutdownFunc func(string) error

func (f OnShutdownFunc) OnShutdown(shutdown string) error {
	return f(shutdown)
}

// ShutdownManager is a manager for shutdown Managers.
type ShutdownManager interface {
	GetName() string
	Start(gs GracefulShutdownIface) error
	ShutdownStart() error
	ShutdownFinish() error
}

type ErrorHandler interface {
	OnError(err error)
}

type ErrorFunc func(err error)

func (f ErrorFunc) OnError(err error) {
	f(err)
}

// GracefulShutdownIface is the interface for graceful shutdown.
type GracefulShutdownIface interface {
	Shutdown(sm ShutdownManager)
	ReportError(err error)
	AddShutdownCallback(callback ShutdownCallback)
	AddShutdownManager(manager ShutdownManager)
	SetErrorHandler(handler ErrorHandler)
}

// GracefulShutdown 结构体：实现了GracefulShutdownIface接口
type GracefulShutdown struct {
	callbacks  []ShutdownCallback
	managers   []ShutdownManager
	errHandler ErrorHandler
}

// AddShutdownCallback implements GracefulShutdownIface.
func (gs *GracefulShutdown) AddShutdownCallback(callback ShutdownCallback) {
	gs.callbacks = append(gs.callbacks, callback)
}

func (gs *GracefulShutdown) AddShutdownManager(sm ShutdownManager) {
	gs.managers = append(gs.managers, sm)
}

func (gs *GracefulShutdown) SetErrorHandler(handler ErrorHandler) {
	gs.errHandler = handler
}

// ReportError implements GracefulShutdownIface.
func (gs *GracefulShutdown) ReportError(err error) {
	if err != nil && gs.errHandler != nil {
		gs.errHandler.OnError(err)
	}
}

// StartShutdown implements GracefulShutdownIface.
func (gs *GracefulShutdown) Shutdown(sm ShutdownManager) {
	gs.ReportError(sm.ShutdownStart())

	var wg sync.WaitGroup
	for _, shutdownCallback := range gs.callbacks {
		wg.Add(1)
		go func(callback ShutdownCallback) {
			defer wg.Done()

			gs.ReportError(callback.OnShutdown(sm.GetName()))
		}(shutdownCallback)
	}

	wg.Wait()

	gs.ReportError(sm.ShutdownFinish())
}

func New() *GracefulShutdown {
	return &GracefulShutdown{
		callbacks: make([]ShutdownCallback, 0),
		managers:  make([]ShutdownManager, 0),
	}
}

func (gs *GracefulShutdown) Start() error {
	for _, manager := range gs.managers {
		if err := manager.Start(gs); err != nil {
			return err
		}
	}
	return nil
}
