package command

import "pokedex/internal/pokecache"

type Config struct {
	Next     string
	Previous string
	Cache    *pokecache.Cache
}
