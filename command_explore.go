package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("location name is required")
	}
	locationName := args[0]

	fmt.Printf("Exploring %s...\n", locationName)

	locationsResp, err := cfg.pokeapiClient.ListLocationArea(locationName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, encounter := range locationsResp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
