package main

import (
	"time"

	"github.com/jverbit/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		commands:      knownCommands(),
		pokeapiClient: pokeClient,
	}

	loopRepl(cfg)
}
