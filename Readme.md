## Go Mux utility to convert Video to GIF

#### This uses ffmpeg lib to covert given video to gif

### Build the Docker image
docker build -t convert-vid-to-gif .

### Run the Docker container
docker run -p 8080:8080 convert-vid-to-gif
