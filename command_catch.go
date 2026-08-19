package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/maualvm/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a valid pokemon name")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	chance := rand.Intn(10)
	caught := true
	if 5 >= chance && chance < 10 {
		caught = false
	}

	if !caught {
		fmt.Printf("%s escaped!\n", args[0])
		return nil
	}

	pokemonDetails, err := cfg.pokeClient.GetPokemonDetails(args[0])
	if err != nil {
		return err
	}

	exists := addToPokedex(pokemonDetails, cfg)
	if !exists {
		fmt.Printf("%s was caught!\n", args[0])
	}
	return nil
}

func addToPokedex(pokemonDetails pokeapi.PokemonResponse, cfg *config) bool {
	savedPokemon := pokeapi.Pokemon{
		ID:     pokemonDetails.ID,
		Name:   pokemonDetails.Name,
		Height: pokemonDetails.Height,
		Weight: pokemonDetails.Weight,
		Stats:  pokemonDetails.Stats,
		Types:  pokemonDetails.Types,
	}

	if _, ok := cfg.pokedex[pokemonDetails.Name]; !ok {
		cfg.pokedex[pokemonDetails.Name] = savedPokemon
		return false
	}

	return true
}
