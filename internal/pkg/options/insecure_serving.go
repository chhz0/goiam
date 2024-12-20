package options

import (
	"github.com/chhz0/goiam/internal/pkg/ginserver"
	"github.com/chhz0/goiam/pkg/simplecobra"
	"github.com/spf13/pflag"
)

const (
	defaultBindAddress = "127.0.0.1"
	defaultBindPort    = 8080
)

type InsecureServingOptions struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	BindPort    int    `json:"bind-port" mapstructure:"bind-port"`
}

// LocalFlagsAndRequired implements simplecobra.Flags.
func (s *InsecureServingOptions) LocalFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	fs = pflag.NewFlagSet("insecure-serving", pflag.ExitOnError)

	fs.StringVar(&s.BindAddress, "insecure.bind-address", s.BindAddress,
		`The IP address on which to serve the --insecure.bind-port
		(set to 0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).`,
	)
	fs.IntVar(&s.BindPort, "insecure.bind-port", s.BindPort,
		`The port on which to serve unsecured, unauthenticated access. It is assumed
		that firewall rules are set up such that this port is not reachable from outside of
		the deployed machine and that port 443 on the iam public address is proxied to this
		port. This is performed by nginx in the default setup. Set to zero to disable.`,
	)

	return
}

// PersistentFlagsAndRequired implements simplecobra.Flags.
func (i *InsecureServingOptions) PersistentFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	return
}

var _ simplecobra.Flags = (*InsecureServingOptions)(nil)

func (s *InsecureServingOptions) AppendToServer(c *ginserver.Config) error {
	c.InsecureServing.BindAddress = s.BindAddress
	c.InsecureServing.BindPort = s.BindPort

	return nil
}

func NewDefaultInsecureOptions() *InsecureServingOptions {
	return &InsecureServingOptions{
		BindAddress: defaultBindAddress,
		BindPort:    defaultBindPort,
	}
}
