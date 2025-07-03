package api

import (
	"encoding/json"
	"pokedex/internal/pokecache"
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
