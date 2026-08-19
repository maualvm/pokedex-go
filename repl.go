package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/maualvm/pokedexcli/internal/pokeapi"
)

const PROMPT = "Pokedex > "

type config struct {
	commands                 map[string]cliCommand
	pokeClient               *pokeapi.Client
	pokedex                  map[string]pokeapi.Pokemon
	nextLocationAreasUrl     *string
	previousLocationAreasUrl *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(PROMPT)
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := words[1:]
		command, exists := cfg.commands[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas in the Pokemon world",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore <location_area>",
			description: "Displays a list of Pokemon located in the specified location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempts to catch a specified Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View details about the specified Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "List out all caught Pokemon",
			callback: commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
