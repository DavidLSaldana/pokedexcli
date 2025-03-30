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

type pokemon struct {
	BaseExperience int `json:"base_experience"`
	Forms          []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Height                 int    `json:"height"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			Order        any `json:"order"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name    string `json:"name"`
	Order   int    `json:"order"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func GetLocationAreaData(url string, cache pokecache.Cache) (locationArea, error) {
	locationAreaData := &locationArea{}
	if url == "" {
		url = apiURL + "location-area"
	} else {
		storedData, ok := cache.Get(url)
		if ok {
			err := json.Unmarshal(storedData, locationAreaData)
			if err != nil {
				return *locationAreaData, errors.New("Error on Unmarshalling location area data from cache")
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

func GetExploreAreaData(area string, cache pokecache.Cache) (exploreArea, error) {
	exploreAreaData := &exploreArea{}

	url := apiURL + "location-area/" + area

	if storedData, ok := cache.Get(url); ok {
		err := json.Unmarshal(storedData, exploreAreaData)
		if err != nil {
			return *exploreAreaData, errors.New("Error on unmarshalling explore area data from cache")
		}
		return *exploreAreaData, nil

	}

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

func GetPokemonData(pokemonName string, cache pokecache.Cache) (pokemon, error) {
	pokemonData := &pokemon{}

	url := apiURL + "Pokemon/" + pokemonName

	if storedData, ok := cache.Get(url); ok {
		err := json.Unmarshal(storedData, pokemonData)
		if err != nil {
			return *pokemonData, errors.New("Error on unmarshalling explore area data from cache")
		}
		return *pokemonData, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return *pokemonData, errors.New("Error on New Request")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return *pokemonData, errors.New("Error on Do Request")
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(pokemonData)
	if err != nil {
		return *pokemonData, errors.New("Error with decoder")
	}

	return *pokemonData, nil

}
