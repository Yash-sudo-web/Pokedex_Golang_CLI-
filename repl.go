package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(*config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		print("Pokédex>> ")
		scanner.Scan()

		words := textClean(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Invalid command. Type 'help' to see the list of commands.")
			continue
		}

	}
}

func textClean(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

type commands struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]commands {
	return map[string]commands{
		"help": {
			name:        "help",
			description: "List all available commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit Pokédex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show the next 20 regions",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous 20 regions",
			callback:    commandMapb,
		},
	}
}
