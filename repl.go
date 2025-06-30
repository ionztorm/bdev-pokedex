package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     string
	previous string
}

func getCommandRegistry() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Provides help information",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "get map locations",
			callback:    commandMap,
		},
	}
}

func runCommand(cfg *config, command string) error {
	if cmd, exists := getCommandRegistry()[command]; exists {
		err := cmd.callback(cfg)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Unknown command:", command)
	}
	return nil
}

func cleanInput(text string) []string {
	return strings.Fields(strings.TrimSpace(strings.ToLower(text)))
}

func startRepl() {

	cfg := &config{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		cleanedInput := cleanInput(input)

		if len(cleanedInput) == 0 {
			continue
		}

		command := cleanedInput[0]

		err := runCommand(cfg, command)
		if err != nil {
			fmt.Println("There was a problem running your command:", err)
		}
	}
}
