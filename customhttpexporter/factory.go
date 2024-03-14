package customhttpexporter

import (
	"context"

	"go.opentelemetry.io/collector/exporter/exporterhelper"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

const (
	// typeStr is the type of the exporter
	typeStr = "customhttpexporter"
)

// NewFactory creates a Datadog exporter factory
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		typeStr,
		createDefaultConfig,
		// Uncomment the exporter type that you would like, change the second parameter as you like. Available options
		//are listed below:
		// component.StabilityLevelUndefined
		// component.StabilityLevelUnmaintained
		// component.StabilityLevelDeprecated
		// component.StabilityLevelDevelopment
		// component.StabilityLevelAlpha
		// component.StabilityLevelBeta
		// component.StabilityLevelStable
		exporter.WithMetrics(createMetricsExporter, component.StabilityLevelBeta),
	)
}

func createDefaultConfig() component.Config {

	return &config{}
}

// createMetricsExporter creates a metrics exporter based on this config.
func createMetricsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Metrics, error) {

	return exporterhelper.NewMetricsExporter(ctx, set, cfg,
		pushMetrics,
		//	The parameters below are optional. Uncomment any as you need.
		//	exporterhelper.WithStart(start component.StartFunc),
		//exporterhelper.WithShutdown(shutdown component.ShutdownFunc),
		//exporterhelper.WithTimeout(timeoutSettings TimeoutSettings),
		//exporterhelper.WithRetry(retrySettings RetrySettings),
		//exporterhelper.WithQueue(queueSettings QueueSettings),
		//exporterhelper.WithCapabilities(capabilities consumer.Capabilities)
	)

}
