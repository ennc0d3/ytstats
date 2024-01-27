package main

import (
	"log"
	"os"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"

	"go.opentelemetry.io/otel/exporters/prometheus"

	"go.opentelemetry.io/otel/sdk/trace"

	"go.opentelemetry.io/otel/sdk/metric"
	//"go.opentelemetry.io/otel/sdk/resource"
	//semconv "go.opentelemetry.io/otel/semconv/v1.17.0"

	"github.com/ennc0d3/yt-stats/internal/api"
)

func main() {
	// Configure zerolog as the global logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := log.New(os.Stderr, "", 0)
	logger = log.New(zerolog.ConsoleWriter{Out: os.Stderr}, "", 0)

	// Initialize OpenTelemetry tracing
	err := initTracingExporter()
	if err != nil {
	}

	// Initialize Prometheus metrics
	err = initMetricExporter()
	if err != nil {
		logger.Fatal("failed to initialize Prometheus metrics")
	}

	// Start the server
	api.StartServer()
}

func initTracingExporter() error {
	// Create a new stdout exporter for tracing
	traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return err
	}

	// Create a new trace provider with the exporter and a simple span processor
	tp := trace.NewTracerProvider(
		trace.WithSyncer(traceExporter),
		trace.WithSampler(trace.AlwaysSample()),
	)

	// Set the global trace provider
	otel.SetTracerProvider(tp)

	return nil
}

func initMetricExporter() error {
	// Create a new Prometheus exporter for metrics

	metricExporter, err := prometheus.New()
	if err != nil {
		return err
	}

	provider := metric.NewMeterProvider(metric.WithReader(metricExporter))
	provider.Meter("youtube-info-api")

	return nil
}
