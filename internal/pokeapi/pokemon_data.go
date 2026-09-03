package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jverbit/pokedexcli/internal/pokecache"
)

func (c *Client) GetPokemonData(name string) (PokemonStats, error) {
	url := baseURL + pokemondata + "/" + name

	if value, ok := c.cache.Get(url); ok {
		pokemonData := PokemonStats{}
		err := json.Unmarshal(value, &pokemonData)
		if err != nil {
			return PokemonStats{}, err
		}
		return pokemonData, nil
	} else {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokemonStats{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonStats{}, err
		}
		defer resp.Body.Close()

		dat, err := io.ReadAll(resp.Body)
		if err != nil {
			return PokemonStats{}, err
		}

		pokecache.AppCache().Add(url, dat)
		pokemonData := PokemonStats{}
		err = json.Unmarshal(dat, &pokemonData)
		if err != nil {
			return PokemonStats{}, err
		}

		return pokemonData, nil
	}
}
