package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	pokemonData, err := cfg.pokeapiClient.GetPokemonData(args[0])
	if err != nil {
		fmt.Println(err)
		return errors.New("error finding pokemon data")
	}
	if len(pokemonData.Name) == 0 {
		fmt.Println("No data was found on", args[0])
		return nil
	}

	randInt := rand.IntN(800)
	if randInt > pokemonData.BaseExperience {
		fmt.Println(args[0], "was caught!")
		cfg.pokedex[pokemonData.Name] = pokemonData
		fmt.Println("You may now inspect", args[0], "with the insepct command")
	} else {
		fmt.Println(args[0], "escaped!")
	}

	return nil
}
