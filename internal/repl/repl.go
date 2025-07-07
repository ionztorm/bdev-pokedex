// Package repl provides a Read-Eval-Print Loop (REPL) interface
// for interacting with the Pokedex application via the command line.
package repl

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/command"
	"pokedex/internal/pokecache"
	"pokedex/internal/utils"
	"time"
)

func StartRepl() {

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

		cleanedInput := utils.CleanInput(input)

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
