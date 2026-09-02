package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jverbit/pokedexcli/internal/pokecache"
)

func (c *Client) ListPokemon(name string) (RespPokemonInLocation, error) {
	url := baseURL + locationAreaEndpoint + "/" + name

	if value, ok := c.cache.Get(url); ok {
		pokemonsResp := RespPokemonInLocation{}
		err := json.Unmarshal(value, &pokemonsResp)
		if err != nil {
			return RespPokemonInLocation{}, err
		}
		return pokemonsResp, nil
	} else {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespPokemonInLocation{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespPokemonInLocation{}, err
		}
		defer resp.Body.Close()

		dat, err := io.ReadAll(resp.Body)
		if err != nil {
			return RespPokemonInLocation{}, err
		}

		pokecache.AppCache().Add(url, dat)
		pokemonsResp := RespPokemonInLocation{}
		err = json.Unmarshal(dat, &pokemonsResp)
		if err != nil {
			return RespPokemonInLocation{}, err
		}

		return pokemonsResp, nil
	}
}
