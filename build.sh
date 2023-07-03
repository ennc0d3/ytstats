#!/bin/bash

# Build the Go binary
go build -o yt-stats ./cmd/...

# Build the Docker image
docker build -t yt-stats:v1 -f docker/utube-stats.dockerfile . 

