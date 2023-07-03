package main

import (
	"log"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"

	"go.opentelemetry.io/otel/exporters/prometheus"

	"go.opentelemetry.io/otel/sdk/trace"

	"go.opentelemetry.io/otel/sdk/metric"
	//"go.opentelemetry.io/otel/sdk/resource"
        //semconv "go.opentelemetry.io/otel/semconv/v1.17.0"


	"github.com/ennc0d3/utube-stats/internal/api"
)

func main() {
	// Configure zerolog as the global logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	//log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Initialize OpenTelemetry tracing
	err := initTracingExporter()
	if err != nil {
		log.Fatal("Failed to initialize tracing exporter")
	}

	// Initialize Prometheus metrics
	err = initMetricExporter()
	if err != nil {
		// zerolog log.Fatal().Err(err).Msg("Failed to initialize metric exporter")
		log.Fatal("Failed to initialize metric exporter")
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

