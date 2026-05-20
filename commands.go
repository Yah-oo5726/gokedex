package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	pokecache "github.com/Yah-oo5726/gokedex/internal"
)

type config struct {
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type map_response struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var commands map[string]cliCommand
var cache *pokecache.Cache

func init() {
	cache = pokecache.NewCache(10 * time.Second)
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
		"map": {
			name:        "map",
			description: "Displays 20 areas. Each subsequent call shows the next 20 locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays 20 areas. Each subsequent call shows the previous 20 locations.",
			callback:    commandMapb,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	output := "Welcome to the Pokedex!\nUsage:\n\n"
	for _, command := range commands {
		output += fmt.Sprintf("%s: %s\n", command.name, command.description)
	}
	fmt.Print(output)
	return nil
}

func commandMap(cfg *config) error {
	if cfg.next == "" {
		fmt.Println("you're on the last page")
		return nil
	}
	cfg.previous, cfg.next = printLocationAreasOnPage(cfg.next)
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	cfg.previous, cfg.next = printLocationAreasOnPage(cfg.previous)
	return nil
}

func printLocationAreasOnPage(url string) (string, string) {
	location_areas, err := cache.GetFrom(url)
	if err != nil {
		fmt.Println("Error getting location areas")
		return "", ""
	}

	result := map_response{}
	err = json.Unmarshal(location_areas, &result)
	if err != nil {
		fmt.Println("Error unmarshaling location areas")
		return "", ""
	}
	for _, location := range result.Results {
		fmt.Println(location.Name)
	}
	return result.Previous, result.Next
}
