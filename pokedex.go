package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Pokemon in Pokédex: ")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Println(pokemon.Name)
	}

	return nil
}
