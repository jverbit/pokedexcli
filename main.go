package main

import (
	"time"

	"github.com/jverbit/pokedexcli/internal/pokeapi"
	"github.com/jverbit/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, &pokecache.Cache{})
	cfg := &config{
		commands:      knownCommands(),
		pokeapiClient: *pokeClient,
	}

	loopRepl(cfg)
}
