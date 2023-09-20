package api

import (
	"net/http"

	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

var (
	serverPort = 8998
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/video-info", handleVideoInfo).Methods(http.MethodGet)
	// Setup other routes...
}

func StartServer() {
	router := mux.NewRouter()
	router.Use(otelmux.Middleware("youtube-info-api"))
	SetupRoutes(router)
	router.Use(hlog.NewHandler(log.Logger))

	// Register Prometheus metrics endpoint
	router.Path("/metrics").Handler(promhttp.Handler())

	// Start the server
	log.Info().Msgf("Starting server on %d", serverPort)
	addr := fmt.Sprintf(":%d", serverPort)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
