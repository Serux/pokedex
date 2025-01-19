package commands

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/Serux/pokedex/apiresponses"
	"github.com/Serux/pokedex/internal/pokecache"
)

type ConfigPoke struct {
	Next     *string
	Previous *string
	Comm     map[string]CliCommand
	Cache    pokecache.PokeCache
	Pokedex  map[string]apiresponses.Pokemon
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*ConfigPoke, []string) error
}

func CommandInspect(c *ConfigPoke, params []string) error {

	Pok, ok := c.Pokedex[params[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name: ", Pok.Name)
	fmt.Println("Height: ", Pok.Height)
	fmt.Println("Weigth: ", Pok.Weight)
	fmt.Println("Stats: ")
	for _, v := range Pok.Stats {
		fmt.Printf(" -%v: %v\n", v.Stat.Name, v.Base_Stat)
	}
	fmt.Println("Types: ")
	for _, v := range Pok.Types {
		fmt.Printf(" - %v\n", v.Type.Name)
	}

	return nil
}

func CommandPokedex(c *ConfigPoke, params []string) error {
	for _, v := range c.Pokedex {
		fmt.Printf(" - %v\n", v.Name)
	}
	return nil
}

func CommandCatch(c *ConfigPoke, params []string) error {
	//TODO fix when no second param.
	if len(params[0]) == 0 {
		return nil
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + params[0]
	fmt.Printf("Throwing a Pokeball at %v...\n", params[0])
	Pok := Get[apiresponses.Pokemon](url, c)
	if rand.Intn(255) > Pok.Base_experience {
		fmt.Printf("%v was caught!\n", Pok.Name)
		c.Pokedex[Pok.Name] = Pok
	} else {
		fmt.Printf("%v escaped!\n", Pok.Name)
	}
	return nil
}

func CommandExplore(c *ConfigPoke, params []string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + params[0]
	fmt.Println("Exploring", params[0], "...")
	explorearea := Get[apiresponses.ExploreArea](url, c)
	fmt.Println("Found Pokemon:")
	for _, v := range explorearea.Pokemon_encounters {
		fmt.Println(" - ", v.Pokemon.Name)
	}

	return nil
}

func CommandMap(c *ConfigPoke, params []string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if c.Next != nil {
		url = *c.Next
	}

	locations := Get[apiresponses.Locationarea](url, c)

	for _, v := range locations.Results {
		fmt.Println(v.Name)
	}
	c.Next = locations.Next
	c.Previous = locations.Previous

	return nil
}

func CommandMapBack(c *ConfigPoke, params []string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if c.Previous != nil {
		url = *c.Previous
	}
	locations := Get[apiresponses.Locationarea](url, c)

	for _, v := range locations.Results {
		fmt.Println(v.Name)
	}
	c.Next = locations.Next
	c.Previous = locations.Previous

	return nil
}

func CommandHelp(c *ConfigPoke, params []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range c.Comm {
		fmt.Printf("%v: %v\n", command.Name, command.Description)
	}
	return nil
}
func CommandExit(_ *ConfigPoke, params []string) error {
	if _, err := fmt.Println("Closing the Pokedex... Goodbye!"); err != nil {
		return err
	}
	os.Exit(0)
	return fmt.Errorf("Cannot close aplication")

}
