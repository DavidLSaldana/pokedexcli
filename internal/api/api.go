package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

var apiURL string = "https://pokeapi.co/api/v2/"

type locationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreaData(url string) (locationArea, error) {
	locationAreaData := &locationArea{}
	if url == "" {
		url = apiURL + "location-area"
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return *locationAreaData, errors.New("Error on NewRequest")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return *locationAreaData, errors.New("Error on DoRequest")
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(locationAreaData)
	if err != nil {
		return *locationAreaData, errors.New("Error with decoder")
	}

	return *locationAreaData, nil

}
