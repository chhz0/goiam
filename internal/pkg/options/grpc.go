package options

import (
	"github.com/chhz0/goiam/pkg/simplecobra"
	"github.com/spf13/pflag"
)

type GRPCOptions struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	BindPort    int    `json:"bind-port" mapstructure:"bind-port"`
	MaxMsgSize  int    `json:"max-msg-size" mapstructure:"max-msg-size"`
}

// LocalFlagsAndRequired implements simplecobra.Flags.
func (g *GRPCOptions) LocalFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	fs = pflag.NewFlagSet("grpc", pflag.ExitOnError)

	fs.StringVar(&g.BindAddress, "grpc.bind-address", g.BindAddress, ""+
		"The IP address on which to serve the --grpc.bind-port(set to 0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).")

	fs.IntVar(&g.BindPort, "grpc.bind-port", g.BindPort, ""+
		"The port on which to serve unsecured, unauthenticated grpc access. It is assumed "+
		"that firewall rules are set up such that this port is not reachable from outside of "+
		"the deployed machine and that port 443 on the iam public address is proxied to this "+
		"port. This is performed by nginx in the default setup. Set to zero to disable.")

	fs.IntVar(&g.MaxMsgSize, "grpc.max-msg-size", g.MaxMsgSize, "gRPC max message size.")

	return
}

// PersistentFlagsAndRequired implements simplecobra.Flags.
func (g *GRPCOptions) PersistentFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	return
}

var _ simplecobra.Flags = (*GRPCOptions)(nil)

func NewDefaultGRPCOptions() *GRPCOptions {
	return &GRPCOptions{
		BindAddress: "127.0.0.1",
		BindPort:    8081,
		MaxMsgSize:  1024 * 1024 * 4,
	}
}
