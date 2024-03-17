package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, params []string) error {
	if len(params) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := params[0]
	fmt.Println("pokemon", name)
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s", pokemon.Name)
	fmt.Printf("Height: %v", pokemon.Height)
	fmt.Printf("Weight: %v", pokemon.Weight)
	fmt.Println("Stats:")
	stats := pokemon.Stats
	for _, stat := range stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}

	return nil
}
