package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
			command_handler()
		}
	}
}
