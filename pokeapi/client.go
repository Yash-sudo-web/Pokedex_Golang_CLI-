package pokeapi

import (
	"net/http"
	"time"

	"github.com/Yash-sudo-web/pokedex/caching"
)

type Client struct {
	httpClient http.Client
	cache      caching.Cache
}

const baseURL = "https://pokeapi.co/api/v2"

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache:      caching.NewCache(cacheInterval),
		httpClient: http.Client{Timeout: timeout},
	}
}
