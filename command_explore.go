package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a valid location area")
	}

	locationArea, err := cfg.pokeClient.GetLocationAreaDetails(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[0])
	pokemonEncounters := locationArea.PokemonEncounters
	pokemonNames := []string{}
	for _, encounter := range pokemonEncounters {
		pokemonNames = append(pokemonNames, encounter.Pokemon.Name)
	}

	fmt.Println("Found Pokemon:")
	for _, pokemonName := range pokemonNames {
		fmt.Printf("- %s\n", pokemonName)
	}

	return nil
}
