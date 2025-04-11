package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ModestMeowth/boot.dev-pokedexcli/internal/pokeapi"
	"github.com/ModestMeowth/boot.dev-pokedexcli/internal/pokecache"
)

type config struct {
    pokeapiClient pokeapi.Client
    pokeapiCache pokecache.Cache
    nextLocationsURL *string
    prevLocationsURL *string
}

type cliCommand struct {
    name string
    description string
    callback func(*config, ...string) error
}

func startRepl(cfg *config) {
    reader := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")
        reader.Scan()

        words := cleanInput(reader.Text())
        if len(words) == 0 {
            continue
        }

        commandName := words[0]

        command, exists := getCommands()[commandName]
        if exists {
            err := command.callback(cfg, words[1:]...)
            if err != nil {
                fmt.Println(err)
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

func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
        "help": {
            name: "help",
            description: "Displays a help message",
            callback: commandHelp,
        },
        "explore": {
            name: "explore",
            description: "List the pokemans at a location",
            callback: commandExplore,
        },
        "map": {
            name: "map",
            description: "Get the next page of locations",
            callback: commandMapF,
        },
        "mapb": {
            name: "map",
            description: "Get the previous page of locations",
            callback: commandMapB,
        },
        "exit": {
            name: "exit",
            description: "Exit the pokedex",
            callback: commandExit,
        },
    }
}
