package main

import (
	"bufio"
	"errors"
	"fmt"
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

	//in progress
	if len(args) > 2 {
		return errors.New("too many args for explore command")
	}
	exploreLocation := args[1]
	exploreData, err := api.GetExploreAreaData(exploreLocation)
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
