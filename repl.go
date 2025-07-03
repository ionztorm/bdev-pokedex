package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/command"
	"pokedex/internal/pokecache"
	"time"
)

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)
	cfg := &command.Config{
		Cache:   pokecache.NewCache(60 * time.Second),
		Pokedex: make(map[string]command.PokemonResp),
	}

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		cleanedInput := cleanInput(input)

		if len(cleanedInput) == 0 {
			continue
		}

		cmd := cleanedInput[0]
		cfg.Args = cleanedInput[1:]

		err := command.RunCommand(cfg, cmd)
		if err != nil {
			fmt.Println("There was a problem running your command:", err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading input: ", err)
	}
}
