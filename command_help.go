package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range cfg.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
