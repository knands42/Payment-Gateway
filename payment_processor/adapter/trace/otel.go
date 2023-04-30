package trace

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"log"
	"os"
)

type OpenTelemetry struct {
	ServiceName      string
	ServiceVersion   string
	ExporterEndpoint string
}

func NewOpenTelemetry(serviceName, serviceVersion, exporterEndpoint string) *OpenTelemetry {
	return &OpenTelemetry{
		ServiceName:      serviceName,
		ServiceVersion:   serviceVersion,
		ExporterEndpoint: exporterEndpoint,
	}
}

func (ot *OpenTelemetry) GetTracer() trace.Tracer {
	logger := log.New(os.Stdout, "zipkin-example", log.Ldate|log.Ltime|log.Llongfile)

	exporter, err := zipkin.New(
		ot.ExporterEndpoint,
		zipkin.WithLogger(logger),
	)

	if err != nil {
		log.Fatal(err)
	}

	batcher := sdktrace.NewBatchSpanProcessor(exporter)

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
