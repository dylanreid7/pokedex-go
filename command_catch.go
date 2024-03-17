package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, params []string) error {
	pokemon := params[0]
	pokemonResp, err := cfg.pokeapiClient.CatchPokemon(pokemon)
	if err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at " + pokemon + "...")
	baseExperience := pokemonResp.BaseExperience
	catchResult := didCatch(baseExperience)
	if catchResult {
		fmt.Println(pokemon + " was caught!")
		cfg.caughtPokemon[pokemon] = pokemonResp
		fmt.Println("You may now inspect it with the inspect command.")
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
		// linear func where baseExp = 36 -> threshold = 0.0, baseExp = 200 -> threshold = .80
		threshold = baseExperienceFlt*(0.8/164) - 0.17561
	} else if baseExperience < 300 {
		// linear func where baseExp = 200 -> threshold = .8, baseExp = 300 -> threshold = .9
		threshold = baseExperienceFlt*(0.1/100) + 0.6
	} else if baseExperience < 400.0 {
		// linear func where baseExp = 300 -> threshold = .9, baseExp - 400 -> threshold = .99
		threshold = baseExperienceFlt*(0.09/100.0) + 0.63
	} else {
		threshold = 0.995
	}
	return randomValue > threshold
}
