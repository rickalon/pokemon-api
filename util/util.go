package util

import (
	"bayau/data"
	"encoding/json"
	"net/http"
)

func WriteJsonError(w http.ResponseWriter, err error, errorCode int) {
	w.WriteHeader(errorCode)
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(&data.ErrorMsg{ErrorMsg: err.Error()})
}

func WriteJsonPokemonData(w http.ResponseWriter, pokemons []*data.Pokemon) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

func WriteJsonPokemonArray(w http.ResponseWriter, pokemons []*data.Pokemon) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}
