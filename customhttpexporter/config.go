package customhttpexporter

import (
	"go.opentelemetry.io/collector/component"
)

// Config defines configuration for your exporter.
type config struct {

	// The URL to send traces to. If omitted the Endpoint + "/v1/traces" will be used.
	Endpoint string `mapstructure:"endpoint"`
}

var _ component.Config = (*config)(nil)

// Validate the configuration for errors to implement the configvalidator interface.
// You can skip this if you do not want to validate your config
func (c *config) Validate() error {
	return nil

}
