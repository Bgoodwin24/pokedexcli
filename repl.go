package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Bgoodwin24/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    *pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			fmt.Println("No input detected, please input your search")
			continue
		}

		commandName := text[0]
		args := []string{}
		if len(text) > 1 {
			args = text[1:]
		}

		cmd, exists := getCommands()[commandName]
		if exists {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"pokedex": {
			name:        "pokedex",
			description: "Displays captured Pokemon",
			callback:    commandPokedex,
		},
		"inspect": {
			name:        "inspect",
			description: "Display stats and types of a captured Pokemon",
			callback:    commandInspect,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch an encountered Pokemon",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore",
			description: "Displays Pokemon in the area",
			callback:    commandExplore,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 areas",
			callback:    commandMapb,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 areas",
			callback:    commandMapf,
		},
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
	}
}
