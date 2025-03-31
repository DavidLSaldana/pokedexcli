package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/DavidLSaldana/pokedexcli/internal/api"
	"github.com/DavidLSaldana/pokedexcli/internal/pokecache"
)

type config struct {
	nextURL string
	prevURL string
	cache   pokecache.Cache
	pokedex map[string]api.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	commandList := getCommandList()
	for _, cmd := range commandList {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config, args []string) error {

	locationArea, err := api.GetLocationAreaData(cfg.nextURL, cfg.cache)
	if err != nil {
		return err
	}

	cfg.nextURL = locationArea.Next
	cfg.prevURL = locationArea.Previous

	for _, location := range locationArea.Results {
		println(location.Name)
	}

	return nil
}

func commandMapB(cfg *config, args []string) error {

	if cfg.prevURL == "" {
		return errors.New("you're on the first page")
	}

	locationArea, err := api.GetLocationAreaData(cfg.prevURL, cfg.cache)
	if err != nil {
		return err
	}

	cfg.nextURL = locationArea.Next
	cfg.prevURL = locationArea.Previous

	for _, location := range locationArea.Results {
		println(location.Name)
	}

	return nil
}

func commandExplore(cfg *config, args []string) error {
	if len(args) > 2 {
		return errors.New("too many args for explore command")
	}
	exploreLocation := args[1]
	exploreData, err := api.GetExploreAreaData(exploreLocation, cfg.cache)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", exploreLocation)
	fmt.Printf("Found Pokemon:\n")
	for _, encounters := range exploreData.PokemonEncounters {
		fmt.Printf(" - %s\n", encounters.Pokemon.Name)
	}

	return nil

}

func commandCatch(cfg *config, args []string) error {
	if len(args) > 2 {
		return errors.New("too many args for catch command")
	}
	pokemonName := args[1]
	pokemonData, err := api.GetPokemonData(pokemonName, cfg.cache)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	isCaught := catchAttempt(pokemonData.BaseExperience)
	if !isCaught {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemonName)
	if _, ok := cfg.pokedex[pokemonName]; ok {
		fmt.Printf("%s has already been caught...letting %s go..\n", pokemonName, pokemonName)
		return nil
	}
	cfg.pokedex[pokemonName] = pokemonData

	return nil
}

func commandInspect(cfg *config, args []string) error {
	if len(args) < 2 {
		return errors.New("not enough args for inspect command")
	}
	if len(args) > 2 {
		return errors.New("too many args for inspect command")
	}
	if len(args) == 2 {
		pokemon, ok := cfg.pokedex[args[1]]
		if !ok {
			return errors.New("that pokemon hasn't been caught yet")
		}
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types")
		for _, types := range pokemon.Types {
			fmt.Printf(" - %s\n", types.Type.Name)
		}
	}

	return nil

}

func commandPokedex(cfg *config, args []string) error {
	if len(args) > 1 {
		return errors.New("too many args for pokedex command")
	}
	if len(cfg.pokedex) == 0 {
		return errors.New("you haven't captured any pokemon yet")
	}
	fmt.Println("List of pokemon...")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil

}

func getCommandList() map[string]cliCommand {
	var commandList = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location areas",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Displays list of all the pokemon at a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "view information on already captured pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "lists the pokemon you have captured",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	return commandList
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{}
	cfg.nextURL = ""
	cfg.prevURL = ""
	cfg.cache = pokecache.NewCache(5 * time.Second)
	cfg.pokedex = make(map[string]api.Pokemon)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(strings.ToLower(scanner.Text()))
		command, ok := getCommandList()[input[0]]
		if !ok {
			fmt.Println("unknown command")
			continue

		}
		err := command.callback(cfg, input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}

}

func cleanInput(text string) []string {
	strSlice := strings.Fields(text)

	return strSlice
}

func catchAttempt(baseExp int) bool {

	testRand := rand.Float64()
	catchRate := ((testRand) * float64(baseExp))

	return float64(100)/float64(catchRate) > float64(0.5)

}
