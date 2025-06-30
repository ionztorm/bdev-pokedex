package main

import (
	"encoding/json"
	"fmt"
)

type locationAreaResp struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

func commandMap(cfg *config) error {
	data, err := getRawData("location-area/")
	if err != nil {
		return fmt.Errorf("error getting location data: %w", err)
	}

	var resp locationAreaResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	cfg.next = resp.Next
	cfg.previous = resp.Previous

	return nil
}
