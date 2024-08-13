package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	videoURL := "https://youtu.be/eVli-tstM5E"

	// Define the output format for yt-dlp using a placeholder for the video title
	outputFile := "../api/videos/%(title)s.%(ext)s"

	// Create the videos directory if it doesn't exist
	if err := os.MkdirAll("videos", os.ModePerm); err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	cmd := exec.Command("yt-dlp", "-f", "bestvideo+bestaudio", "--merge-output-format", "mp4", "-o", outputFile, videoURL)

	// Run the command and capture the output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error downloading video: %v", err)
	}

	fmt.Println("Video downloaded and merged successfully.")
}
