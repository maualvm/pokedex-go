package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("you have no pokemon")
	}

	fmt.Println("Your Pokedex:")
	for pokemonName := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemonName)
	}
	return nil
}
