package main

import (
	"fmt"
)

func commandInspect(cfg *config, name ...string) error {
	if pokemon, ok := cfg.pokedex[name[0]]; ok {
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)
		fmt.Println("Stats:")
		for _, subset := range pokemon.Stats {
			fmt.Printf("  - %s: %d\n", subset.Stat.Name, subset.BaseStat)
		}
		fmt.Println("Types:")
		for _, subset := range pokemon.Types {
			fmt.Printf("  - %s\n", subset.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}
