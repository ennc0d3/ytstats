package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/hlog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

func handleVideoInfo(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("youtube-info-api")
	meter := otel.Meter("youtube-info-api")
	//TODO: Initialize once
	videoViewsCounter, _ := meter.Float64Counter(
		"utube_video_views",
		metric.WithDescription("Total number of video views"))

	// Start a new span for the handler
	_, span := tracer.Start(r.Context(), "handleVideoInfo")
	defer span.End()

	videoID := r.URL.Query().Get("videoid")
	if videoID == "" {
		hlog.FromRequest(r).Error().Msg("videoid parameter is missing")
		http.Error(w, "videoid parameter is missing", http.StatusBadRequest)
		return
	}

	// Perform the desired logging
	hlog.FromRequest(r).Info().Msg("Retrieving video information")

	// Call the YouTube API to retrieve video statistics
	statistics, err := GetVideoStatistics(videoID)
	if err != nil {
		hlog.FromRequest(r).Error().Err(err).Msg("failed to retrieve video statistics")
		http.Error(w, "failed to retrieve video statistics", http.StatusInternalServerError)
		return
	}

	// Increment the video views counter
	videoViewsCounter.Add(context.Background(), 1)

	// ...
	hlog.FromRequest(r).Info().Msgf("stats: %v", *statistics)

	// Convert the response struct to JSON
	jsonResponse, err := json.Marshal(statistics.rawStats)
	if err != nil {
		hlog.FromRequest(r).Error().Err(err).Msg("failed to marshal JSON response")
		http.Error(w, "failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
