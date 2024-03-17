package main

import "fmt"

func commandExplore(cfg *config, params []string) error {
	locationResp, err := cfg.pokeapiClient.ExploreLocation(params[0])
	if err != nil {
		return err
	}

	fmt.Println("Exploring " + params[0] + " area...")
	fmt.Println("Found Pokemon:")
	// fmt.Println(locationResp)

	// locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	// if err != nil {
	// 	return err
	// }
	pokemonEncounters := locationResp.PokemonEncounters

	for _, encounter := range pokemonEncounters {
		fmt.Println("- " + encounter.Pokemon.Name)
	}
	// cfg.nextLocationsURL = locationsResp.Next
	// cfg.prevLocationsURL = locationsResp.Previous

	// for _, loc := range locationsResp.Results {
	// 	fmt.Println(loc.Name)
	// }
	return nil
}
