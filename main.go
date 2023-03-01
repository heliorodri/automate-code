package main

import "os"

func main() {
	apiKey := getApiKey()

	if apiKey != "" {
		println("done")
	}

}

func getApiKey() string {
	apiKey := os.Getenv("AUTOMATE_CODE_API_KEY")

	if apiKey == "" {
		panic("Missing API key")
	}

	return apiKey
}
