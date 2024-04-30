package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("No location provided")
	}
	name := args[0]
	location, err := cfg.pokeapiClient.ListPokemon(name)
	if err != nil {
		return errors.New("Location does not exist")
	}
	fmt.Println("Exporing " + location.Name)
	fmt.Println("Pokemon Found:")
	for _, pokemon := range location.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
