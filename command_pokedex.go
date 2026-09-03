package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("You haven't caught a Pokemon yet!")
	} else {
		for _, pokemon := range cfg.pokedex {
			fmt.Printf(" - %s\n", pokemon.Name)
		}
	}
	return nil
}
