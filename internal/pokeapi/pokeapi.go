package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/maualvm/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

type locationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type locationAreaResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type PokemonResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

type Pokemon struct {
	ID     int
	Name   string
	Height int
	Weight int
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

const POKEAPI_URL = "https://pokeapi.co/api/v2"

func NewClient(cacheInterval time.Duration) *Client {
	return &Client{
		httpClient: http.Client{},
		cache:      *pokecache.NewCache(cacheInterval),
	}
}

func (c *Client) ListLocationAreas(pageUrl *string) (locationAreasResponse, error) {
	url := POKEAPI_URL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	locationAreas := locationAreasResponse{}
	data, err := c.makeGetRequest(url)
	if err != nil {
		return locationAreasResponse{}, err
	}

	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return locationAreasResponse{}, err
	}

	return locationAreas, nil
}

func (c *Client) GetLocationAreaDetails(locationArea string) (locationAreaResponse, error) {
	url := POKEAPI_URL + "/location-area/" + locationArea

	locationAreaDetails := locationAreaResponse{}
	data, err := c.makeGetRequest(url)
	if err != nil {
		return locationAreaResponse{}, err
	}

	if err := json.Unmarshal(data, &locationAreaDetails); err != nil {
		return locationAreaResponse{}, err
	}

	return locationAreaDetails, nil
}

func (c *Client) GetPokemonDetails(pokemonName string) (PokemonResponse, error) {
	url := POKEAPI_URL + "/pokemon/" + pokemonName

	pokemonDetails := PokemonResponse{}
	data, err := c.makeGetRequest(url)
	if err != nil {
		return PokemonResponse{}, err
	}

	if err := json.Unmarshal(data, &pokemonDetails); err != nil {
		return PokemonResponse{}, err
	}

	return pokemonDetails, nil
}

func (c *Client) makeGetRequest(url string) ([]byte, error) {
	cacheRes, found := c.cache.Get(url)
	if found {
		return cacheRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	c.cache.Add(url, data)
	return data, nil
}
