package main

import (
	"fmt"
	"io"
	"net/http"
	"pokedex/utils"
)

const baseURL = "https://pokeapi.co/api/v2/"

func getRawData(endpoint string) ([]byte, error) {
	fullURL := baseURL + endpoint

	fmt.Println("Debug: fullURL is: ", fullURL)

	req, err := http.NewRequest("GET", fullURL, nil)
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
