package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	test()
}

func test() {
	url := "http://localhost:8080/convert-to-gif"

	// Create a sample video file to upload (replace with your actual video file)
	videoFilePath := "./video.mp4"
	videoFile, err := os.Open(videoFilePath)
	if err != nil {
		fmt.Println("Error opening video file:", err)
		return
	}
	defer videoFile.Close()

	// Create a buffer to store the request body
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Create the form file field
	fileWriter, err := writer.CreateFormFile("file", "video.mp4")
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}

	// Copy the file content to the form file field
	_, err = io.Copy(fileWriter, videoFile)
	if err != nil {
		fmt.Println("Error copying file content:", err)
		return
	}

	// Close the multipart writer
	writer.Close()

	// Create the HTTP request with the given body and Content-Type
	request, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Content-Type header with the boundary parameter
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// Make the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	// Extract the base name of the video file
	baseName := filepath.Base(videoFilePath)
	// Find the last occurrence of dot
	lastDotIndex := strings.LastIndex(baseName, ".")

	if lastDotIndex != -1 {
		// Split the string based on the last dot
		baseName = baseName[:lastDotIndex]
		fmt.Println("Base Name:", baseName)
	} else {
		// No dot found, handle the case accordingly
		fmt.Println("No dot found in the string")
	}

	// Save the response to a file
	responseFilePath := baseName + ".gif"

	// Create or open a file to save the response
	file, err := os.Create(responseFilePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Copy the response body to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error saving response to file:", err)
		return
	}

	fmt.Println("API response saved to:", responseFilePath)
}
