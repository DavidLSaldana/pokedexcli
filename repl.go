package main

import (
	"bufio"
	//"errors"
	"fmt"
	"os"
	"strings"
)

type config struct {
	nextURL string
	prevURL string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	commandList := getCommandList()
	for _, cmd := range commandList {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	//get location-area endpoint and print

	return nil
}

func commandMapB(cfg *config) error {
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
			callback:    commandMap,
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
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(strings.ToLower(scanner.Text()))
		command, ok := getCommandList()[input[0]]
		if !ok {
			fmt.Println("unknown command")
			continue

		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
	}

}

func cleanInput(text string) []string {
	strSlice := strings.Fields(text)

	return strSlice
}
