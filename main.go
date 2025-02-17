package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"net/url"
)

// OMDbAPIResponse struct to hold response from OMDb API
type OMDbAPIResponse struct {
	Search []struct {
		Title   string `json:"Title"`
		Year    string `json:"Year"`
		Genre   string `json:"Genre"`
		IMDBID  string `json:"imdbID"`
		Rating  string `json:"imdbRating"`
		Type    string `json:"Type"`
	} `json:"Search"`
	Response string `json:"Response"`
}

func main() {
	// Replace with your own OMDb API Key
	apiKey := os.Getenv("OMDP_KEY")

	// Define the search query URL
	baseURL := "http://www.omdbapi.com/?apikey=" + apiKey
	searchURL := baseURL + "&t=Longlegs&type=movie"

	// Send HTTP GET request to OMDb API
	resp, err := http.Get(searchURL)
	if err != nil {
		log.Fatalf("Error fetching data from OMDb API: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Parse the JSON response
	var apiResponse OMDbAPIResponse
	fmt.Println(body)
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Fatalf("Error unmarshaling response: %v", err)
	}

	// Check if the response contains movies
	if apiResponse.Response == "False" {
		fmt.Println("No movies found.")
		return
	}

	fmt.Println(apiResponse)

	// Display horror movies and their ratings
	fmt.Println("Current Horror Movies and Ratings:")
	for _, movie := range apiResponse.Search {
		fmt.Printf("Title: %s\nYear: %s\nRating: %s\n\n", movie.Title, movie.Year, movie.Rating)
	}
}
