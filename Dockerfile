# Use the official Golang image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Install ffmpeg
RUN apt-get update && \
    apt-get install -y ffmpeg

# Copy the local code to the container
COPY main.go .
COPY go.mod .
COPY go.sum .

# Download Go dependencies
RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
