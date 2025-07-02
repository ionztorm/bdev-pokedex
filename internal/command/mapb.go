package command

import (
	"fmt"
	"pokedex/internal/api"
)

func cmapb(cfg *Config) error {
	fullURL := api.BaseURL + "location-area"

	if cfg.Previous != "" {
		fullURL = cfg.Previous
	}

	var resp locationAreaResp
	if err := Fetch(cfg, fullURL, &resp); err != nil {
		return fmt.Errorf("failed to fetch location data: %w", err)
	}

	if resp.Previous == "" {
		fmt.Println("\nyou're on the first page")
		fmt.Println()
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
