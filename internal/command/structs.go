package command

import "pokedex/internal/pokecache"

type Config struct {
	Next     string
	Previous string
	Cache    *pokecache.Cache
	Args     []string
}

type locationAreaResp struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type locationPokemonResp struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
