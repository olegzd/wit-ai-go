package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	apiToken := os.Getenv("WIT_TOKEN")
	if len(apiToken) <= 0 {
		log.Fatal("No api token provided")
	}
	witAiSpeechURL := "https://api.wit.ai/speech"

	request, err := http.NewRequest("POST", witAiSpeechURL, nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	request.Header.Set("Authorization", "Bearer "+apiToken)
	request.Header.Set("Content-Type", "audio/wav")
}
