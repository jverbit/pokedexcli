package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jverbit/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

type config struct {
	commands         map[string]cliCommand
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.PokemonStats
}

func loopRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := cleanInput(scanner.Text())
			if len(input) == 0 {
				continue
			}

			if command, exists := knownCommands()[input[0]]; exists {
				if len(input[1:]) == 0 && input[0] == "explore" {
					fmt.Println("Explore command needs a location")
					continue
				}
				if len(input[1:]) == 0 && input[0] == "catch" {
					fmt.Println("Catch command needs a target")
					continue
				}
				command.callback(cfg, input[1:]...)
			} else {
				fmt.Println("Unknown command")
				continue
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			continue
		}

	}
}

func knownCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays names of 20 location areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays names of the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore [location-area]",
			description: "Lists all of the Pokemon in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch [pokemon]",
			description: "Attempts to catch a Pokemon",
			callback:    commandCatch,
		},
	}
}
