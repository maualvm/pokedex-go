package main

import (
	"time"

	"github.com/maualvm/pokedexcli/internal/pokeapi"
)

const (
	cacheInterval = 5 * time.Second
)

func main() {
	pokeClient := pokeapi.NewClient(cacheInterval)
	pokedex := make(map[string]pokeapi.Pokemon)
	cfg := &config{
		commands:   getCommands(),
		pokeClient: pokeClient,
		pokedex:    pokedex,
	}
	startRepl(cfg)
}
