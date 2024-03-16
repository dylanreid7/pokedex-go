package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg config) error {
	locations, err := ListLocations(cfg.next)
	if err != nil {
		return err
	}

	setConfig(locations.Next, locations.Previous)

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandBMap(cfg config) error {
	if cfg.prev == nil {
		return errors.New("you're on the first page")
	}
	locations, err := ListLocations(cfg.next)
	if err != nil {
		return err
	}

	setConfig(locations.Next, locations.Previous)

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func ListLocations(url string) (Location, error) {
	defaultUrl := "https://pokeapi.co/api/v2/location-area"
	if url == "" {
		url = defaultUrl
	}
	resp, err := http.Get(url)
	if err != nil {
		return Location{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locations := Location{}
	err = json.Unmarshal(body, &locations)

	if err != nil {
		return Location{}, err
	}

	return locations, nil
}

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
