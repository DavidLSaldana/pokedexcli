package api

import (
	"encoding/json"
)

var apiURL string = "https://pokeapi.co/api/v2/"

type locationArea struct {
	count    int    `json:"count"`
	next     string `json:"next"`
	previous string `json:"previous"`
	results  []struct {
		name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getLocationArea() {

}
