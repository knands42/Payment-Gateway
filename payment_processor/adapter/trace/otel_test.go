package trace

import (
	"context"
	"github.com/caiofernandes00/payment-gateway/adapter/trace/exporter"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/baggage"
	"testing"
	"time"
)

func Test_TracingAFunction(t *testing.T) {
	t.Run("should trace a function using zipkin as the exporter", func(t *testing.T) {
		// Arrange
		ctx := baggage.ContextWithoutBaggage(context.Background())
		zipkin := exporter.NewZipkinExporter("").GetExporter()
		ot := NewOpenTelemetry(zipkin).GetTracer()

		testFn := func() {
			println("test before sleep")
			time.Sleep(1 * time.Second)
			println("test after sleep")
		}

		// Act
		TraceFn(ot, ctx, "", testFn)

		// Assert
		require.NotEmpty(t, ot)
		require.NotEmpty(t, ctx)
		require.NotEmpty(t, zipkin)
	})
}
