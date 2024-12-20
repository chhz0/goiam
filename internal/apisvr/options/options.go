package options

import (
	"github.com/chhz0/goiam/internal/pkg/options"
	"github.com/chhz0/goiam/pkg/simplecobra"
	"github.com/spf13/pflag"
)

type APIServerOptions struct {
	Server          *options.ServerRunOptions       `json:"server" mapstructure:"server"`
	SecureServing   *options.SecureServingOptions   `json:"secure" mapstructure:"secure"`
	InsecureServing *options.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	Fearure         *options.FeatureOptions         `json:"feature" mapstructure:"feature"`
	MySQL           *options.MySQLOptions           `json:"mysql" mapstructure:"mysql"`
}

// LocalFlagsAndRequired implements simplecobra.Flags.
func (o *APIServerOptions) LocalFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	fs = pflag.NewFlagSet("api-server", pflag.ExitOnError)

	serverFlags, _ := o.Server.LocalFlagsAndRequired()
	secureServingFlags, _ := o.SecureServing.LocalFlagsAndRequired()
	insecureServingFlags, _ := o.InsecureServing.LocalFlagsAndRequired()
	featureFlags, _ := o.Fearure.LocalFlagsAndRequired()
	mysqlFlags, _ := o.MySQL.LocalFlagsAndRequired()

	fs.AddFlagSet(serverFlags)
	fs.AddFlagSet(secureServingFlags)
	fs.AddFlagSet(insecureServingFlags)
	fs.AddFlagSet(featureFlags)
	fs.AddFlagSet(mysqlFlags)

	return
}

// PersistentFlagsAndRequired implements simplecobra.Flags.
func (o *APIServerOptions) PersistentFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	return
}

var _ simplecobra.Flags = (*APIServerOptions)(nil)

func NewNilAPIOptions() *APIServerOptions {
	return &APIServerOptions{
		Server:          options.NewDefaultServerRunOptions(),
		SecureServing:   options.NewDefaultSecureOptions(),
		InsecureServing: options.NewDefaultInsecureOptions(),
		Fearure:         options.NewDefaultFeatureOptions(),
		MySQL:           options.NewDefaultMySQLOptions(),
	}
}