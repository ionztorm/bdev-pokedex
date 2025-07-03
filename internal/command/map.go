package command

import (
	"fmt"
	"pokedex/internal/api"
)

func cmap(cfg *Config) error {
	fullURL := api.BaseURL + "location-area"
	if cfg.Next != "" {
		fullURL = cfg.Next
	}

	var resp locationAreaResp
	if err := Fetch(cfg, fullURL, &resp); err != nil {
		return fmt.Errorf("failed to fetch location data: %w", err)
	}

	if resp.Next == "" {
		fmt.Println("\nYou're on the last page")
	}

	fmt.Println()
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println()

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	return nil
}
