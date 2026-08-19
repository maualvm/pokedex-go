package main

import (
	"errors"
	"fmt"
)

func commandMapB(cfg *config, args ...string) error {
	if cfg.nextLocationAreasUrl == nil {
		return errors.New("you're on the first page")
	}

	locationAreas, err := cfg.pokeClient.ListLocationAreas(cfg.previousLocationAreasUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationAreasUrl = locationAreas.Next
	cfg.previousLocationAreasUrl = locationAreas.Previous

	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)
	}
	return nil
}
