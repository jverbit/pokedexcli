package pokeapi

type RespPokemonInLocation struct {
	Name               string `json:"name"`
	Pokemon_encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
