package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

const pokemonURL = "https://pokeapi.co/api/v2/pokemon/"

type pokemon struct {
	LocationURL string `json:"location_area_encounters"`
	Name        string `json:"name"`
}

type location struct {
	Locations map[string]string `json:"location_area"`
}

type pokemonOutput struct {
	Name      string
	Locations []string
}

func getLocations(locationURL string) ([]string, error) {
	httpResponse, err := http.Get(locationURL)
	if err != nil {
		return nil, err
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var decodedLocations []location
	err = json.Unmarshal(bodyContent, &decodedLocations)
	if err != nil {
		return nil, err
	}

	var locations []string
	for _, key := range decodedLocations {
		locations = append(locations, key.Locations["name"])
	}

	return locations, nil
}

func getPokemonDetails(idOrName string) (string, []string, error) {
	httpResponse, err := http.Get(pokemonURL + idOrName)
	if err != nil {
		return "", nil, err
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return "", nil, err
	}

	var pokemonDetails pokemon
	err = json.Unmarshal(bodyContent, &pokemonDetails)
	if err != nil {
		return "", nil, err
	}

	locations, err := getLocations(pokemonDetails.LocationURL)
	if err != nil {
		return "", nil, err
	}
	return pokemonDetails.Name, locations, nil
}

func main() {
	pokemonIdPtr := flag.String("pokemon", "1", "name or ordinal number of the pokemon")
	flag.Parse()

	name, locations, err := getPokemonDetails(*pokemonIdPtr)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "this pokemon might not exist"),
		)
	}

	pokemonJson, err := json.Marshal(pokemonOutput{name, locations})

	fmt.Printf("%s\n", pokemonJson)
}
