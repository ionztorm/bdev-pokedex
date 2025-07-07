// Package api provides utilities for fetching and caching data
// from external HTTP APIs, with JSON unmarshaling support.
package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedex/internal/pokecache"
	"pokedex/internal/utils"
)

func Fetch(cache pokecache.Cacher, url string, target any) error {
	data, exists := cache.Get(url)
	if !exists {
		var err error
		data, err = getRawDataFromURL(url)
		if err != nil {
			return err
		}
		cache.Add(url, data)
	}
	return json.Unmarshal(data, target)
}

func getRawDataFromURL(endpoint string) ([]byte, error) {

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("arror fetching data: %w", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("an error occurred: %w", err)
	}

	defer utils.SafeClose(res.Body)

	if res.StatusCode != http.StatusOK {
		fmt.Println(res.StatusCode)
		return nil, fmt.Errorf("problem with response")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return body, nil
}
