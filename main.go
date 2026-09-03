package main

import (
	"time"

	"github.com/jverbit/pokedexcli/internal/pokeapi"
	"github.com/jverbit/pokedexcli/internal/pokecache"
)

func main() {
	sharedCache := pokecache.AppCache()
	pokeClient := pokeapi.NewClient(5*time.Second, sharedCache)
	cfg := &config{
		commands:      knownCommands(),
		pokedex:       map[string]pokeapi.PokemonStats{},
		pokeapiClient: *pokeClient,
	}

	loopRepl(cfg)
}
