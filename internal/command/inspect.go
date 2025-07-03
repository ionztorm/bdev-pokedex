package command

import "fmt"

func cinspect(cfg *Config) error {
	pokemon := cfg.Args[0]
	if pokemon == "" {
		fmt.Println("the 'inspect' command requires a pokemon name, for example: 'inspect pikachu'")
		return nil
	}

	result, exists := cfg.Pokedex[pokemon]
	if !exists {
		fmt.Printf("\nYou have not caught %s yet!\n\n", pokemon)
		return nil
	}

	// name, height, weight, stats and type(s) of the Pokemon.
	fmt.Println()
	fmt.Printf("Name: %s\n", result.Name)
	fmt.Printf("Height: %d\n", result.Height)
	fmt.Printf("Weight: %d\n", result.Weight)

	fmt.Println("Stats:")
	for _, stat := range result.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range result.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}
	fmt.Println()

	return nil

}
