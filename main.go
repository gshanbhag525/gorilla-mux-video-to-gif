package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/convert-to-gif", func(w http.ResponseWriter, r *http.Request) {
		// Parse the form data to get the file
		err := r.ParseMultipartForm(10 << 20) // 10 MB limit
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, _, err := r.FormFile("video")
		if err != nil {
			http.Error(w, "Error retrieving the file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Save the uploaded file
		uploadPath := "./uploads/"
		err = os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			http.Error(w, "Error creating upload directory", http.StatusInternalServerError)
			return
		}

		fileName := "uploaded_video.mp4"
		filePath := uploadPath + fileName
		out, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Error creating the file", http.StatusInternalServerError)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			http.Error(w, "Error saving the file", http.StatusInternalServerError)
			return
		}

		// Convert the uploaded video to GIF
		outputGifPath := "./output.gif"
		err = convertVideoToGif(filePath, outputGifPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Video successfully converted to GIF.")
	}).Methods("POST")

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
