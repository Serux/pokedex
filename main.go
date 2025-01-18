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

var commands map[string]cliCommand

func main() {
	commands = map[string]cliCommand{
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

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userText := scanner.Text()
		userCommands := cleanInput(userText)

		if len(userCommands) == 0 {
			continue
		}
		if _, ok := commands[userCommands[0]]; !ok {
			continue
		}

		commands[userCommands[0]].callback()

	}
}
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}
func commandExit() error {
	if _, err := fmt.Println("Closing the Pokedex... Goodbye!"); err != nil {
		return err
	}
	os.Exit(0)
	return fmt.Errorf("Cannot close aplication")

}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)

}
