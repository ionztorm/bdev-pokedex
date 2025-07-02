package command

import (
	"encoding/json"
	"pokedex/internal/api"
)

func Fetch(cfg *Config, url string, v any) error {
	data, exists := cfg.Cache.Get(url)
	if !exists {
		var err error
		data, err = api.GetRawDataFromURL(url)
		if err != nil {
			return err
		}
		cfg.Cache.Add(url, data)
	}
	return json.Unmarshal(data, v)
}
