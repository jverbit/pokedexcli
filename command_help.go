package main

import "fmt"

func commandHelp() error {
	fmt.Printf("\n")
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	for _, cmd := range knownCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Printf("\n")
	return nil
}
