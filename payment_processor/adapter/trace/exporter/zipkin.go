package exporter

import (
	"go.opentelemetry.io/otel/exporters/zipkin"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"log"
	"os"
)

type ZipkinExporter struct {
	ExporterEndpoint string
}

func NewZipkinExporter(exporterEndpoint string) *ZipkinExporter {
	return &ZipkinExporter{
		ExporterEndpoint: exporterEndpoint,
	}
}

func (z *ZipkinExporter) GetExporter() sdktrace.SpanExporter {
	logger := log.New(os.Stdout, "zipkin-example", log.Ldate|log.Ltime|log.Llongfile)

	exporter, err := zipkin.New(
		z.ExporterEndpoint,
		zipkin.WithLogger(logger),
	)

	if err != nil {
		log.Fatal(err)
	}

	return exporter
}
