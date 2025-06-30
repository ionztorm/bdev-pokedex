package command

import (
	"encoding/json"
	"fmt"
	"pokedex/api"
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

	data, err := api.GetRawDataFromURL(fullURL)
	if err != nil {
		return fmt.Errorf("error getting location data: %w", err)
	}

	var resp locationAreaResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
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
