package api

import (
	"encoding/json"
	"errors"
	"github.com/DavidLSaldana/pokedexcli/internal/pokecache"
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

type exploreArea struct {
	ID       int `json:"id"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func GetLocationAreaData(url string, cache pokecache.Cache) (locationArea, error) {
	locationAreaData := &locationArea{}
	if url == "" {
		url = apiURL + "location-area"
	} else {
		stored, ok := cache.Get(url)
		if ok {
			err := json.Unmarshal(stored, locationAreaData)
			if err != nil {
				return *locationAreaData, errors.New("Error on Unmarshaling from cache")
			}
			return *locationAreaData, nil

		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return *locationAreaData, errors.New("Error on New Request")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return *locationAreaData, errors.New("Error on Do Request")
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(locationAreaData)
	if err != nil {
		return *locationAreaData, errors.New("Error with decoder")
	}

	return *locationAreaData, nil
}

func GetExploreAreaData(area string) (exploreArea, error) {
	exploreAreaData := &exploreArea{}

	url := apiURL + "location-area/" + area

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return *exploreAreaData, errors.New("Error on New Request")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return *exploreAreaData, errors.New("Error on Do Request")
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(exploreAreaData)
	if err != nil {
		return *exploreAreaData, errors.New("Error with decoder")
	}

	return *exploreAreaData, nil

}
