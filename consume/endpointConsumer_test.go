package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Response struct {
	Name string `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}


type Pokemon struct {
	EntryNo int `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species""`
}

type PokemonSpecies struct {
	Name string `json:"name"`
	//Url string `json:"url"`
}

func TestGetEndpoint(t *testing.T) {
	response, err := GetEndpoint("http://pokeapi.co/api/v2/pokedex/hoenn/")

	if err != nil {
		t.Error("Failed with error", err)
		return
	}

	var responseObj Response
	jsonErr := json.Unmarshal([]byte(response), &responseObj)

	if jsonErr != nil {
		t.Error("Failed with error", err)
		return
	}

	fmt.Printf("%+v\n", responseObj)

}