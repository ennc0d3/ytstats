package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	



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
	log.Info().Msg("Starting server on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}

