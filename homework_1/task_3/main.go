package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/sethgrid/pester"
)

type pokemon struct {
	Location_area_encounters	string
	Name						string
}

type location struct {
	Location_area				map[string]string
}


func linearBackoff(retry int) time.Duration {
	return time.Duration(retry) * time.Second
}

func getLocations(locationURL string) ([]string, error) {
	httpClient := pester.New()
	httpClient.Backoff = linearBackoff

	httpResponse, err := httpClient.Get(locationURL)
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
		locations = append(locations, key.Location_area["name"])
	}

	return locations, nil
}

func getPokemonDetails(idOrName string) (string, []string, error) {
	pokemonURL := "https://pokeapi.co/api/v2/pokemon/" + idOrName

	httpClient := pester.New()
	httpClient.Backoff = linearBackoff

	httpResponse, err := httpClient.Get(pokemonURL)
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

	locations, err := getLocations(pokemonDetails.Location_area_encounters)
	if err != nil {
		return "", nil, err
	}
	return pokemonDetails.Name, locations, nil
}


func main() {
	var pokemonId string

	fmt.Printf("Please, enter the name or the ordinal number of a pokemon: ")
	fmt.Scanf("%v", &pokemonId)

	name, locations, err := getPokemonDetails(pokemonId)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "this pokemon might not exist"),
		)
	}

	fmt.Printf("The pokemon '%s' can be found here: %v\n", name, strings.Join(locations, ", "))
}
