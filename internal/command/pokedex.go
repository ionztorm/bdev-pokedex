package command

import "fmt"

func cpokedex(cfg *Config) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty!")
		return nil
	}

	fmt.Println()
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.Pokedex {
		fmt.Printf(" - %v\n", pokemon.Name)
	}
	fmt.Println()

	return nil
}
