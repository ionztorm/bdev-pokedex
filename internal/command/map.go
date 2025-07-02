package command

import (
	"fmt"
	"pokedex/internal/api"
)

type locationAreaResp struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

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
