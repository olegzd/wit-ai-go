package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	MsgID string `json:"msg_id"`
	Text  string `json:"_text"`
}

func main() {
	apiToken := os.Getenv("WIT_TOKEN")
	if len(apiToken) <= 0 {
		log.Fatal("No api token provided")
	}
	witAiSpeechURL := "https://api.wit.ai/speech"

	sampleFile, err := os.Open("sample.wav")
	if err != nil {
		log.Fatal("Error while opening sample file: ", err)
	}

	request, err := http.NewRequest("POST", witAiSpeechURL, sampleFile)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	request.Header.Set("Authorization", "Bearer "+apiToken)
	request.Header.Set("Content-Type", "audio/wav")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		fmt.Printf("Error while sending request!")
		log.Fatal(error)
	}

	var sample = Response{}

	json.NewDecoder(response.Body).Decode(&sample)

	fmt.Printf("Result ID: %s\n", sample.MsgID)
	fmt.Printf("Result _text: %s\n", sample.Text)
}
