package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Serux/pokedex/apiresponses"
	"github.com/Serux/pokedex/commands"
	"github.com/Serux/pokedex/internal/pokecache"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	config := commands.ConfigPoke{
		Comm: map[string]commands.CliCommand{
			"exit": {
				Name:        "exit",
				Description: "Exit the Pokedex",
				Callback:    commands.CommandExit,
			},
			"help": {
				Name:        "help",
				Description: "Displays a help message",
				Callback:    commands.CommandHelp,
			},
			"map": {
				Name:        "map",
				Description: "Shows next 20 locations-areas",
				Callback:    commands.CommandMap,
			},
			"mapb": {
				Name:        "mapb",
				Description: "Shows previous 20 locations-areas",
				Callback:    commands.CommandMapBack,
			},
			"explore": {
				Name:        "explore",
				Description: "Shows pokemons of a location",
				Callback:    commands.CommandExplore,
			},
			"catch": {
				Name:        "catch",
				Description: "Tryes to catch a pokemon",
				Callback:    commands.CommandCatch,
			},
			"inspect": {
				Name:        "inspect",
				Description: "Inspect a catched pokemon",
				Callback:    commands.CommandInspect,
			},
			"pokedex": {
				Name:        "pokedex",
				Description: "Shows captured pokemons",
				Callback:    commands.CommandPokedex,
			},
		},
		Cache:   pokecache.NewCache(time.Second * 5),
		Pokedex: map[string]apiresponses.Pokemon{},
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userText := scanner.Text()
		userCommands := cleanInput(userText)

		if len(userCommands) == 0 {
			continue
		}
		if _, ok := config.Comm[userCommands[0]]; !ok {
			continue
		}
		params := []string{}
		if len(userCommands) > 1 {
			params = userCommands[1:]
		}

		config.Comm[userCommands[0]].Callback(&config, params)

	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)

}
