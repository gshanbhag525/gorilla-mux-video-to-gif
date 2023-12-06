# 🎥➡️🖼️ Gorilla Mux Video to GIF Converter 🎥➡️🖼️

This project utilizes the Gorilla Mux router for handling HTTP requests to convert video files to GIFs using FFmpeg. The application provides a simple API endpoint to upload a video file, process it with FFmpeg, and return the resulting GIF. The application is written in Golang

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Docker Container](#docker-container)
- [Contributing](#contributing)
- [License](#license)


## 🛠️ Installation 🛠️

1. Clone the repository:

```bash
git clone https://github.com/gshanbhag525/gorilla-mux-video-to-gif.git && 
cd gorilla-mux-video-to-gif
```

2. Install dependencies:

```bash
go mod download
```

3. Build the project:

```bash
go build
```

4. Run the executable:

```bash
./gorilla-mux-video-to-gif
```

## 🚀 Usage 🚀

1. Ensure the application is running.

2. Use your preferred API client (e.g., curl, Postman) to make a POST request to http://localhost:8080/convert-to-gif with the video file attached.

```bash
curl -X POST -F "file=@/path/to/your/video.mp4" http://localhost:8080/convert-to-gif
```

3. The server will process the video file and return a link to download the generated GIF.

## 🐳 Docker Container 🐳

To run the application in a Docker container:

1. Build the Docker image:

```bash
docker build -t gorilla-mux-video-to-gif .
```
2. Run the Docker container:

```bash
docker run -p 8080:8080 gorilla-mux-video-to-gif
```

3. Follow the usage instructions mentioned above.

## 🤝 Contributing 🤝

If you'd like to contribute to this project, please follow the [contribution guidelines](Contributing.md).


## 📝 License 📝

This project is licensed under the [MIT License](LICENSE).