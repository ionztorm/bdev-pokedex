package command

import "fmt"

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
