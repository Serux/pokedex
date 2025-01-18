package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		ret := s.Text()
		cli := cleanInput(ret)
		fmt.Printf("Your command was: %v\n", cli[0])

	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)

}
