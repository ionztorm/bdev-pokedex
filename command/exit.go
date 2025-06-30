package command

import (
	"fmt"
	"os"
)

func cexit(cfg *Config) error {
	fmt.Println()
	fmt.Println("Closing the Pokedex... Goodbye!")
	fmt.Println()
	os.Exit(0)
	return nil
}
