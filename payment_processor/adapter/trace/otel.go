package trace

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

type OpenTelemetryExporter interface {
	GetExporter() sdktrace.SpanExporter
}

type OpenTelemetry struct {
	exporter sdktrace.SpanExporter
}

func NewOpenTelemetry(exporter sdktrace.SpanExporter) *OpenTelemetry {
	return &OpenTelemetry{
		exporter: exporter,
	}
}

func (ot *OpenTelemetry) GetTracer() trace.Tracer {
	batcher := sdktrace.NewBatchSpanProcessor(ot.exporter)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(batcher),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL, semconv.ServiceNameKey.String("payment-processor"),
		)),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	tracer := otel.Tracer("io.opentelemetry.traces.goapp")
	return tracer
}

func TraceFn(otel trace.Tracer, ctx context.Context, tracingName string, fn func()) {
	_, t := otel.Start(ctx, tracingName)
	fn()
	t.End()
}
