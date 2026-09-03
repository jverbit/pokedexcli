package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	locationMap, err := cfg.pokeapiClient.ListPokemon(args[0])
	if err != nil {
		return errors.New("error finding location data")
	}
	if len(locationMap.Name) == 0 {
		fmt.Println("No data found on location", args[0])
		return nil
	}
	fmt.Printf("Exploring %s...\n", locationMap.Name)
	fmt.Println("Found Pokemon: ")
	for _, encounter := range locationMap.PokemonEncounters {
		fmt.Println("-", encounter.Pokemon.Name)
	}

	return nil
}
