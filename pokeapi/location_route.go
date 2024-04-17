package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		println("Cache hit")
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	locationResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespShallowLocations{}, err
	}
	println("Cache miss")
	c.cache.Add(url, data)
	return locationResp, nil
}

func (c *Client) ListPokemon(location string) (LocationArea, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		locationResp := LocationArea{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationArea{}, err
		}
		println("Cache hit")
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}
	locationResp := LocationArea{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return LocationArea{}, err
	}
	println("Cache miss")
	c.cache.Add(url, data)
	return locationResp, nil
}

func (c *Client) CatchPokemon(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon

	if val, ok := c.cache.Get(url); ok {
		locationResp := Pokemon{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Pokemon{}, err
		}
		println("Cache hit")
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	locationResp := Pokemon{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return Pokemon{}, err
	}
	println("Cache miss")
	c.cache.Add(url, data)
	return locationResp, nil
}
