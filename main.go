package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/convert-to-gif", func(w http.ResponseWriter, r *http.Request) {
		// Replace "path/to/your/input/video.mp4" with the actual path to your video file
		inputVideoPath := "./video.mp4"
		// Replace "path/to/your/output/output.gif" with the desired output GIF file path
		outputGifPath := "./output.gif"

		err := convertVideoToGif(inputVideoPath, outputGifPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Video successfully converted to GIF.")
	}).Methods("GET")

	http.Handle("/", r)
	fmt.Println("Server is running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func convertVideoToGif(inputVideoPath, outputGifPath string) error {
	// Check if the input video file exists
	if _, err := os.Stat(inputVideoPath); os.IsNotExist(err) {
		return fmt.Errorf("Input video file does not exist")
	}

	// Use ffmpeg to convert the video to gif
	cmd := exec.Command("ffmpeg", "-i", inputVideoPath, "-vf", "fps=10,scale=320:-1", outputGifPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error converting video to GIF: %v", err)
	}

	return nil
}
