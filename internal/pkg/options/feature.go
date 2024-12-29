package options

import (
	"github.com/chhz0/goiam/internal/pkg/ginserver"
	"github.com/chhz0/goiam/pkg/simplecobra"
	"github.com/spf13/pflag"
)

const (
	defaultEnableProfiling = false
	defaultEnableMetrics   = false
)

// FeatureOptions 关联到 server 的功能选项
type FeatureOptions struct {
	EnableProfiling bool `json:"enable-profiling" mapstructure:"enable-profiling"`
	EnableMetrics   bool `json:"enable-metrics" mapstructure:"enable-metrics"`
}

// LocalFlagsAndRequired implements simplecobra.Flags.
func (f *FeatureOptions) LocalFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	fs = pflag.NewFlagSet("feature", pflag.ExitOnError)
	fs.BoolVar(&f.EnableProfiling, "feature.enable-profiling", f.EnableProfiling,
		`Enable the profiling endpoint. Install /debug/pprof router handler.`)
	fs.BoolVar(&f.EnableMetrics, "feature.enable-metrics", f.EnableMetrics,
		`Enable the metrics endpoint. Install /metrics router handler.`)

	return
}

// PersistentFlagsAndRequired implements simplecobra.Flags.
func (f *FeatureOptions) PersistentFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	return
}

var _ simplecobra.Flags = (*FeatureOptions)(nil)

func (f *FeatureOptions) AppendToServer(c *ginserver.Config) error {
	c.EnableMetrics = f.EnableMetrics
	c.EnableProfiling = f.EnableProfiling

	return nil
}

func NewDefaultFeatureOptions() *FeatureOptions {
	return &FeatureOptions{
		EnableProfiling: defaultEnableProfiling,
		EnableMetrics:   defaultEnableMetrics,
	}
}
