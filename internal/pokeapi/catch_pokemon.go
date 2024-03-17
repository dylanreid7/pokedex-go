package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// Catch Pokemon -
func (c *Client) CatchPokemon(pokemon string) (RespShallowPokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespShallowPokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespShallowPokemon{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	locationResp := RespShallowPokemon{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	c.cache.Add(url, dat)
	return locationResp, nil
}
