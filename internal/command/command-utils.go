// Package command implements the CLI command system,
// including the central command registry, individual command logic,
// and shared structs used across commands.
package command

import "fmt"

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

func GetCommandRegistry() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    cexit,
		},
		"help": {
			Name:        "help",
			Description: "Provides help information",
			Callback:    chelp,
		},
		"map": {
			Name:        "map",
			Description: "get first, or next 20 map locations",
			Callback:    cmap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "get the previous 20 map locations",
			Callback:    cmapb,
		},
		"explore": {
			Name:        "explore",
			Description: "See which pokemon can be found in the chosen area. Usage: 'explore <location name>'",
			Callback:    cexplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Attempt to catch a pokemon. Usage: 'catch <pokemon name>'",
			Callback:    ccatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Look up information about a pokemon you've caught. Usage: 'inspect <pokemon name>'",
			Callback:    cinspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "See a list of the pokemon you've caught and have in your Pokedex",
			Callback:    cpokedex,
		},
	}
}

func RunCommand(cfg *Config, cmd string) error {
	if c, exists := GetCommandRegistry()[cmd]; exists {
		err := c.Callback(cfg)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Unknown command:", cmd)
	}
	return nil
}
