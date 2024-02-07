package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	apiToken, ok := os.LookupEnv("API_TOKEN")
	if !ok {
		panic("API_TOKEN environment variable not set")
	}
	url, ok := os.LookupEnv("URL")
	if !ok {
		panic("URL environment variable not set")
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Token %s", apiToken))

	client := &http.Client{}
	for {
		time.Sleep(15 * time.Second)
		response, err := client.Do(request)
		if err != nil {
			log.Print(err)
			continue
		}
		defer response.Body.Close()
		fmt.Println(response.Status)
	}
}
