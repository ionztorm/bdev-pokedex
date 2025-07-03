package command

import (
	"fmt"
	"pokedex/internal/api"
)

func cexplore(cfg *Config) error {
	area := cfg.Args[0]
	if area == "" {
		fmt.Println("the explore command requires a an area location. for example: 'explore sinnoh-pokemon-league-area'")
		return nil
	}

	fullURL := api.BaseLocationAreaURL + area

	var resp locationPokemonResp

	if err := api.Fetch(cfg.Cache, fullURL, &resp); err != nil {
		return fmt.Errorf("error fetching pokemon for %v: %w", area, err)
	}

	fmt.Printf("\nFound the following Pokemon in %s\n\n", area)
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	fmt.Println()

	return nil
}
