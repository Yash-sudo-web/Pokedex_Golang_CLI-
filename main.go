package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Yash-sudo-web/pokedex/pokeapi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
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

	godotenv.Load()

	portstr := os.Getenv("PORT")
	if portstr == "" {
		log.Fatal("PORT not found in .env file")
	}

	dbstr := os.Getenv("DB_URL")
	if dbstr == "" {
		log.Fatal("DB_URL not found in .env file")
	}

	con, err1 := sql.Open("postgres", dbstr)
	if err1 != nil {
		log.Fatal("Cannot connect to the DB.", err1)
	}

	err2 := con.Ping()
	if err2 != nil {
		log.Fatal("Cannot ping the DB.", err2)
	}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portstr,
	}

	go startRepl(cfg)

	fmt.Printf("Server is running on http://localhost:%s\n", portstr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
