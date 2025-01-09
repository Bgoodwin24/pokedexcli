package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandMap map[string]cliCommand

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commandMap = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanText := cleanInput(text)

		if len(cleanText) == 0 {
			fmt.Println("No input detected, please input your search")
			continue
		}

		if cmd, ok := commandMap[cleanText[0]]; ok {
			err := cmd.callback()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for _, cmd := range commandMap {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}
	return words
}
