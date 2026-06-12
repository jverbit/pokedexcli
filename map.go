package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var nextMap string
var prevMap string

type Response struct {
	Count    int
	Next     string
	Previous string
	Results  []Results
}

type Results struct {
	Name string
	Url  string
}

func commandMap() error {
	if len(nextMap) == 0 {
		nextMap = "https://pokeapi.co/api/v2/location-area/"
	}
	getPokeAPI(nextMap)
	return nil
}

func commandMapb() error {
	if prevMap == "" {
		fmt.Printf("you're on the first page\n")
	}
	if len(prevMap) == 0 {
		prevMap = "https://pokeapi.co/api/v2/location-area/"
	}
	getPokeAPI(prevMap)
	return nil
}

func getPokeAPI(mapSegment string) {
	res, err := http.Get(mapSegment)
	if err != nil {
		fmt.Printf("Couldn't reach the Pokemon api.\n")
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading: %v", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return
	}
	/*if err != nil {
		fmt.Printf("Couldn't read the Pokemon server response.\n")
		return
	}*/

	response := Response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("Error parsing json: %v", err)
		return
	}

	nextMap = response.Next
	prevMap = response.Previous

	for _, locations := range response.Results {
		fmt.Printf("%s\n", locations.Name)
	}
}
