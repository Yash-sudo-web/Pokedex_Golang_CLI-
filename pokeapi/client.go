package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

const baseURL = "https://pokeapi.co/api/v2"

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{Timeout: timeout},
	}
}
