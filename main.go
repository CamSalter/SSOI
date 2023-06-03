package main

//import necessary packages
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// define a struct to hold the response data
type ApodResponse struct {
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	Hdurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	Url            string `json:"url"`
}

func main() {
	//send a get request to the NASA API
	resp, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY")
	//error handling
	if err != nil {
		log.Fatal(err)
	}
	//handle closing of the response body
	defer resp.Body.Close()

	apodResponse := &ApodResponse{}
	//decode the JSON response body into the struct
	err = json.NewDecoder(resp.Body).Decode(apodResponse)
	//error handling
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Title: %s\n", apodResponse.Title)
	fmt.Printf("Date: %s\n", apodResponse.Date)
	fmt.Printf("Explanation: %s\n", apodResponse.Explanation)
	fmt.Printf("URL: %s\n", apodResponse.Url)
}
