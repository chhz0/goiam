package options

import (
	"github.com/chhz0/goiam/internal/pkg/ginserver"
	"github.com/chhz0/goiam/pkg/simplecobra"
	"github.com/spf13/pflag"
)

const (
	defaultServerMode    = "release"
	defaultServerHealthz = false
)

type ServerRunOptions struct {
	Mode        string   `json:"mode" mapstructure:"mode"`
	Healthz     bool     `json:"healthz" mapstructure:"healthz"`
	Middlewares []string `json:"middlewares" mapstructure:"middlewares"`
}

// LocalFlagsAndRequired implements simplecobra.Flags.
func (s *ServerRunOptions) LocalFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	fs = pflag.NewFlagSet("server-run", pflag.ExitOnError)

	fs.StringVar(&s.Mode, "server.mode", s.Mode,
		`Start the server in a specific mode. Supported modes: debug, release, test.`)
	fs.BoolVar(&s.Healthz, "server.healthz", s.Healthz,
		`Enable the healthz endpoint. Install /healthz router handler.`)
	fs.StringSliceVar(&s.Middlewares, "server.middlewares", s.Middlewares,
		`List of allowed middlewares for server, comma separated. If this list is empty default middlewares will be used.`)

	return
}

// PersistentFlagsAndRequired implements simplecobra.Flags.
func (s *ServerRunOptions) PersistentFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	return
}

var _ simplecobra.Flags = (*ServerRunOptions)(nil)

func (s *ServerRunOptions) AppendToServer(c *ginserver.Config) error {
	c.Mode = s.Mode
	c.Healthz = s.Healthz
	c.Middlewares = s.Middlewares

	return nil
}

func NewDefaultServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{
		Mode:        defaultServerMode,
		Healthz:     defaultServerHealthz,
		Middlewares: make([]string, 10),
	}
}
