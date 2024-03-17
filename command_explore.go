package main

import "fmt"

func commandExplore(cfg *config, params []string) error {
	locationResp, err := cfg.pokeapiClient.ExploreLocation(params[0])
	if err != nil {
		return err
	}

	fmt.Println("Exploring " + params[0] + " area...")
	fmt.Println("Found Pokemon:")

	pokemonEncounters := locationResp.PokemonEncounters

	for _, encounter := range pokemonEncounters {
		fmt.Println("- " + encounter.Pokemon.Name)
	}

	return nil
}
