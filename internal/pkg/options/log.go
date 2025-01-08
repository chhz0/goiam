package options

// TODO add log options
type LogOptions struct {
	OutputPaths      []string `json:"output-paths" mapstructure:"output-paths"`
	ErrorOutputPaths []string `json:"error-output-paths" mapstructure:"error-output-paths"`
}
