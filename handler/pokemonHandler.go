package handler

import (
	"bayau/data"
	"bayau/util"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

var NUMBER_POKEMONS = 1025
var POKEMONS_REQUESTED = 10

func PokemonHandler(w http.ResponseWriter, r *http.Request) {

	mapRandomPokemon := util.RandomPokemons(POKEMONS_REQUESTED)
	pokemons := make([]*data.Pokemon, POKEMONS_REQUESTED)

	for number, it := range mapRandomPokemon {
		url := "https://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(number)
		resp, err := http.Get(url)
		if err != nil {
			util.WriteJsonError(w, err)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			util.WriteJsonError(w, err)
			return
		}
		err = json.Unmarshal(body, &pokemons[it])
		if err != nil {
			util.WriteJsonError(w, err)
			return
		}
	}
	log.Println(pokemons)
	//util.WriteJsonPokemonArray(w, response)
}
