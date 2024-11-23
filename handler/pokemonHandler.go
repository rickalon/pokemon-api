package handler

import (
	"bayau/data"
	"bayau/util"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

var NUMBER_POKEMONS = 1025
var POKEMONS_REQUESTED = 10

func PokemonHandler(w http.ResponseWriter, r *http.Request) {

	mapPokemon := util.RandomPokemons(POKEMONS_REQUESTED)

	response := make([][]interface{}, POKEMONS_REQUESTED)
	iterator := 0
	for i := range mapPokemon {
		url := "https://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(i)
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
		pokemon := data.NewPokemon()
		err = json.Unmarshal(body, pokemon)
		if err != nil {
			util.WriteJsonError(w, err)
			return
		}
		response[iterator] = []interface{}{pokemon.Name, pokemon.Weight, pokemon.Sprite.ImageURL}
		iterator++
	}
	util.WriteJsonPokemonArray(w, response)
}
