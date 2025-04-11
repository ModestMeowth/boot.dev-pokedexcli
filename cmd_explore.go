package main

import (
    "errors"
    "fmt"
)

func commandExplore(cfg *config, options ...string) error {
    if len(options) == 0 {
        return errors.New("no location provided")
    }

    fmt.Printf("Exploring %s...\n", options[0])

    area, err := cfg.pokeapiClient.GetLocation(options[0])
    if err != nil {
        return err
    }

    fmt.Println("Found Pokemon:")

    for _, pokemon := range area.Encounters {
        fmt.Printf("- %s\n", pokemon.Pokemon.Name)
    }

    return nil
}
