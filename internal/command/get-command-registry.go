package command

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
			Description: "see which pokemon can be found in the chosen area",
			Callback:    cexplore,
		},
	}
}
