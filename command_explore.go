package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	locationMap, err := cfg.pokeapiClient.ListPokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", locationMap.Name)
	fmt.Println("Found Pokemon: ")
	for _, encounter := range locationMap.Pokemon_encounters {
		fmt.Println("-", encounter.Pokemon.Name)
	}

	return nil
}
