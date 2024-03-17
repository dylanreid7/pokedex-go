package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, params []string) error {
	pokemon := cfg.caughtPokemon
	if len(pokemon) == 0 {
		errors.New("You don't have any pokemon")
	}
	fmt.Println("Your Pokedex:")
	for _, poke := range pokemon {
		fmt.Printf(" - %s\n", poke.Name)
	}

	return nil
}
