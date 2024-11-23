package util

import (
	"bayau/data"
	"encoding/json"
	"net/http"
)

var RESULT_ARRAY = 1
var RESULT_OBJECT = 2

func WriteJsonError(w http.ResponseWriter, err error, errorCode int) {
	w.WriteHeader(errorCode)
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(&data.ErrorMsg{ErrorMsg: err.Error()})
}

func WriteJsonPokemon(option int, w http.ResponseWriter, pokemons []*data.Pokemon) {
	switch option {
	case 1:
		writeJsonPokemonArray(w, pokemons)
	case 2:
		writeJsonPokemonData(w, pokemons)
	}
}

func writeJsonPokemonData(w http.ResponseWriter, pokemons []*data.Pokemon) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

func writeJsonPokemonArray(w http.ResponseWriter, pokemons []*data.Pokemon) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")

	json.NewEncoder(w).Encode(pokemonsToArray(pokemons))
}

func pokemonsToArray(pokemons []*data.Pokemon) *[][]interface{} {
	slice := make([][]interface{}, len(pokemons))
	for i, val := range pokemons {
		slice[i] = []interface{}{val.Name, val.Weight, val.Sprite.ImageURL}
	}
	return &slice
}
