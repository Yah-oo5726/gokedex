package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	command_configs := config{
		next:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		previous: "",
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		clean_input := cleanInput(scanner.Text())
		command, exists := commands[clean_input[0]]
		if !exists {
			fmt.Println("Unknown command")
		} else {
			command_handler := command.callback
			command_handler(&command_configs)
		}
		scanner.Err()
	}
}
