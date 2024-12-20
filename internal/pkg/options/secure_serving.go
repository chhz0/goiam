package options

import (
	"github.com/chhz0/goiam/internal/pkg/ginserver"
	"github.com/chhz0/goiam/pkg/simplecobra"
	"github.com/spf13/pflag"
)

const (
	defaultSecureAddress  = "0.0.0.0"
	defaultSecurePort     = 8443
	defaultSecureRequired = true
	defaultSecureCertDir  = "/var/run/iam"
	defaultSecurePairName = "iam"
)

type SecureServingOptions struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	BindPort    int    `json:"bind-port" mapstructure:"bind-port"`
	Required    bool   `json:"required" mapstructure:"required"`
	TLS         *TLS   `json:"tls" mapstructure:"tls"`
}

type TLS struct {
	CertDir  string   `json:"cert-dir" mapstructure:"cert-dir"`
	PairName string   `json:"pair-name" mapstructure:"pair-name"`
	CertKey  *CertKey `json:"cert-key" mapstructure:"cert-key"`
}

type CertKey struct {
	CertFile       string `json:"cert-file" mapstructure:"cert-file"`
	PrivateKeyFile string `json:"private-key-file" mapstructure:"private-key-file"`
}

// LocalFlagsAndRequired implements simplecobra.Flags.
func (s *SecureServingOptions) LocalFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	fs = pflag.NewFlagSet("secure-serving", pflag.ExitOnError)

	fs.StringVar(&s.BindAddress, "secure.bind-address", s.BindAddress,
		`The IP address on which to listen for the --secure.bind-port port.
		The associated interface(s) must be reachable by the rest of the engine, and by CLI/web
		clients. If blank, all interfaces will be used (0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).`)
	desc := "The port on which to serve HTTPS with authentication and authorization."

	if s.Required {
		desc += " It cannot be switched off with 0."
	} else {
		desc += " If 0, don't serve HTTPS at all."
	}

	fs.IntVar(&s.BindPort, "secure.bind-port", s.BindPort, desc)

	fs.StringVar(&s.TLS.CertDir, "secure.tls.cert-dir", s.TLS.CertDir, ""+
		"The directory where the TLS certs are located. "+
		"If --secure.tls.cert-key.cert-file and --secure.tls.cert-key.private-key-file are provided, "+
		"this flag will be ignored.")

	fs.StringVar(&s.TLS.PairName, "secure.tls.pair-name", s.TLS.PairName, ""+
		"The name which will be used with --secure.tls.cert-dir to make a cert and key filenames. "+
		"It becomes <cert-dir>/<pair-name>.crt and <cert-dir>/<pair-name>.key")

	fs.StringVar(&s.TLS.CertKey.CertFile, "secure.tls.cert-key.cert-file", s.TLS.CertKey.CertFile, ""+
		"File containing the default x509 Certificate for HTTPS. (CA cert, if any, concatenated "+
		"after server cert).")

	fs.StringVar(&s.TLS.CertKey.PrivateKeyFile, "secure.tls.cert-key.private-key-file",
		s.TLS.CertKey.PrivateKeyFile, ""+
			"File containing the default x509 private key matching --secure.tls.cert-key.cert-file.")

	return
}

// PersistentFlagsAndRequired implements simplecobra.Flags.
func (s *SecureServingOptions) PersistentFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	return
}

var _ simplecobra.Flags = (*SecureServingOptions)(nil)

func (s *SecureServingOptions) AppendToServer(c *ginserver.Config) error {
	c.SecureServing.BindAddress = s.BindAddress
	c.SecureServing.BindAddress = s.BindAddress
	c.SecureServing.CertKey = ginserver.CertKey{
		Cert: s.TLS.CertKey.CertFile,
		Key:  s.TLS.CertKey.PrivateKeyFile,
	}

	return nil
}

func NewDefaultSecureOptions() *SecureServingOptions {
	return &SecureServingOptions{
		BindAddress: defaultSecureAddress,
		BindPort:    defaultSecurePort,
		Required:    defaultSecureRequired,
		TLS: &TLS{
			CertDir:  defaultSecureCertDir,
			PairName: defaultSecurePairName,
			CertKey: &CertKey{
				CertFile:       "",
				PrivateKeyFile: "",
			},
		},
	}
}
