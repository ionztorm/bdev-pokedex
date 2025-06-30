package command

import (
	"encoding/json"
	"fmt"
	"pokedex/api"
)

func cmapb(cfg *Config) error {
	fullURL := api.BaseURL + "location-area"
	isFirstPage := true

	if cfg.Previous != "" {
		fullURL = cfg.Previous
		isFirstPage = false
	}

	data, err := api.GetRawDataFromURL(fullURL)
	if err != nil {
		return fmt.Errorf("error fetching previous map data: %w", err)
	}

	var resp locationAreaResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return fmt.Errorf("error parsing mapb response: %w", err)
	}

	if isFirstPage {
		fmt.Println()
		fmt.Println("you're on the first page")
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
