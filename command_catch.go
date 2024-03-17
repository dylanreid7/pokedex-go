package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, params []string) error {
	pokemon := params[0]
	locationResp, err := cfg.pokeapiClient.CatchPokemon(pokemon)
	if err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at " + pokemon + "...")
	baseExperience := locationResp.BaseExperience
	catchResult := didCatch(baseExperience)
	if catchResult {
		fmt.Println(pokemon + " was caught!")
	} else {
		fmt.Println(pokemon + " escaped!")
	}
	return nil
}

func didCatch(baseExperience int) bool {
	randomValue := rand.Float32()
	baseExperienceFlt := float32(baseExperience)
	var threshold float32 = 1.0
	if baseExperience < 200 {
		fmt.Println("less than 200")
		// linear func where x = 36 -> y = 0.0, x = 200 -> y = .80
		threshold = baseExperienceFlt*(0.8/164) - 0.17561
	} else if baseExperience < 300 {
		fmt.Println("200 to 300")
		// linear func where x = 200 -> y = .8, x = 300 -> y = .9
		threshold = baseExperienceFlt*(0.1/100) + 0.6
	} else if baseExperience < 400.0 {
		fmt.Println("300 to 400")
		// linear func where x = 300 -> y = .9, x - 400 -> y = .99
		threshold = baseExperienceFlt*(0.09/100.0) + 0.63
	} else {
		// if base exp >= 400
		threshold = 0.995
	}
	return randomValue > threshold
}
