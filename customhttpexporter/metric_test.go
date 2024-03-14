package customhttpexporter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func TestMetricsExporterNoErrors(t *testing.T) {
	f := NewFactory()
	lme, err := f.CreateMetricsExporter(context.Background(), exportertest.NewNopCreateSettings(), f.CreateDefaultConfig())
	require.NotNil(t, lme)
	assert.NoError(t, err)

	assert.NoError(t, lme.ConsumeMetrics(context.Background(), pmetric.NewMetrics()))

	assert.NoError(t, lme.Shutdown(context.Background()))
}
