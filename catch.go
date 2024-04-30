package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("pokemon name not provided")
	}
	name := args[0]
	pokemon, err := cfg.pokeapiClient.CatchPokemon(name)
	if err != nil {
		return errors.New("pokemon does not exist")
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("%s has base experience of %d!\n", pokemon.Name, pokemon.BaseExperience)

	fmt.Printf("Throwing Pokéball at %s...\n", name)

	if res > 50 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)

	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil
}
