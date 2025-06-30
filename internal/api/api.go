package api

import (
	"fmt"
	"io"
	"net/http"
	"pokedex/internal/utils"
)

func GetRawDataFromURL(endpoint string) ([]byte, error) {

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
