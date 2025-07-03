package command

import (
	"fmt"
	"math/rand"
	"pokedex/internal/api"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func ccatch(cfg *Config) error {
	pokemon := cfg.Args[0]

	if pokemon == "" {
		fmt.Println("the 'catch' command requires a pokemon name, for example: 'catch pikachu'")
		return nil
	}

	fullURL := api.BasePokemonURL + pokemon

	var resp PokemonResp

	if err := api.Fetch(cfg.Cache, fullURL, &resp); err != nil {
		return fmt.Errorf("unable to fetch pokemon (%s): %w", pokemon, err)
	}

	fmt.Printf("\nThrowing a Pokeball at %s...\n\n", pokemon)

	chance := catchChance(resp.BaseExperience)
	roll := rng.Float64()

	if roll < chance {
		fmt.Printf("%s was caught!\n\n", pokemon)
		fmt.Printf("You may now inspect it with the inspect command: 'inspect %s'\n\n", pokemon)
		cfg.Pokedex[pokemon] = resp
	} else {
		fmt.Printf("%s escaped!\n\n", pokemon)
	}

	return nil
}

func catchChance(baseExp int) float64 {
	const C = 75.0

	chance := C / (C + float64(baseExp))

	if chance < 0.05 {
		chance = 0.05
	}
	if chance > 0.95 {
		chance = 0.95
	}

	return chance
}
