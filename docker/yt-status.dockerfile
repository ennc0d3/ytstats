# Use a lightweight base image
FROM golang:1.17-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the project source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o yt-stats /app/cmd/...

# Use a minimal base image for the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the previous stage
COPY --from=build /app/yt-stats .

# Expose the port that the API listens on
EXPOSE 8080

# Set the command to run the binary
CMD ["./yt-stats"]

