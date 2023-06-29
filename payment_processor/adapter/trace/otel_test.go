package trace

import (
	"context"
	"testing"
	"time"

	"github.com/caiofernandes00/payment-gateway/adapter/trace/exporter"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/baggage"
)

func Test_TracingAFunction(t *testing.T) {
	t.Run("should trace a function using zipkin as the exporter", func(t *testing.T) {
		// Arrange
		ctx := baggage.ContextWithoutBaggage(context.Background())
		zipkin := exporter.NewZipkinExporter("").GetExporter()
		ot := NewOpenTelemetry(zipkin)
		trace := ot.TraceFn(ot.GetTracer())
		anyValue := "any value"

		testFn := func(ctx context.Context) {
			println("test before sleep")
			time.Sleep(1 * time.Second)
			println("test after sleep")
			anyValue = "new value"
		}

		// Act
		newCtx := trace(ctx, "", testFn)

		// Assert
		require.NotEqual(t, ctx, newCtx)
		require.Equal(t, "new value", anyValue)
	})
}
