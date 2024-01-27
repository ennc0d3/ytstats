package api

import (
	"os"

	"context"

	"google.golang.org/api/option"
	youtube "google.golang.org/api/youtube/v3"
)

type VideoStatistics struct {
	rawStats *youtube.VideoStatistics
	// Add other statistics fields as needed
}

//google.golang.org/api/youtube/v3#VideoStatistics

func GetVideoStatistics(videoID string) (*VideoStatistics, error) {

	ctx := context.TODO()
	apiKey := os.Getenv("YTSTATS_API_KEY")

	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	part := []string{"statistics"}

	videoListCall := youtubeService.Videos.List(part).Id(videoID)
	videoListResp, err := videoListCall.Do()
	if err != nil {
		return nil, err
	}

	rawStats := videoListResp.Items[0].Statistics
	statistics := &VideoStatistics{rawStats: rawStats}

	return statistics, nil
}
