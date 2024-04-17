package main

import (
	"time"

	"github.com/Yash-sudo-web/pokedex/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)
}
