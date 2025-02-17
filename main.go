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
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalf("Error parsing response: %v", err)
	}

	for key, value := range data {
		fmt.Printf("%s: %v\n", key, value)
	}

