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
	callback    func() error
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
	}
}

func runCommand(command string) error {
	if cmd, exists := getCommandRegistry()[command]; exists {
		err := cmd.callback()
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

		err := runCommand(command)
		if err != nil {
			fmt.Println("There was a problem running your command:", err)
		}
	}
}
