package main

import (
	"fmt"
)

func commandMapF(cfg *config, args ...string) error {
	locationAreas, err := cfg.pokeClient.ListLocationAreas(cfg.nextLocationAreasUrl)
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
