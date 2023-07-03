package api

import (
	"os"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
	"net/http"
)

type VideoStatistics struct {
	Views uint64
	// Add other statistics fields as needed
}

func GetVideoStatistics(videoID string) (*VideoStatistics, error) {
	youtubeService, err := youtube.New(&http.Client{
		Transport: &transport.APIKey{Key: os.Getenv("YOUTUBE_API_KEY")},
	})
	if err != nil {
		return nil, err
	}

	kind := []string{"statistics"}

	videoCall := youtubeService.Videos.List(kind).Id(videoID)
	videoResponse, err := videoCall.Do()
	if err != nil {
		return nil, err
	}

	statistics := &VideoStatistics{
		Views: videoResponse.Items[0].Statistics.ViewCount,
		// Extract other statistics fields as needed
	}

	return statistics, nil
}

