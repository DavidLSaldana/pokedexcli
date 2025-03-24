package main

import (
	"bufio"
	//"errors"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	commandList := getCommandList()
	for _, cmd := range commandList {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap() error {
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
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(strings.ToLower(scanner.Text()))
		command, ok := getCommandList()[input[0]]
		if !ok {
			fmt.Println("unknown command")
			continue

		}
		err := command.callback()
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
	}

}

func cleanInput(text string) []string {
	strSlice := strings.Fields(text)

	return strSlice
}
