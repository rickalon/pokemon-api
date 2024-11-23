package util

import (
	"bayau/data"
	"encoding/json"
	"net/http"
)

type ErrorMsg struct {
	erroMsg string `json:"error"`
}

func WriteJsonError(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(&ErrorMsg{err.Error()})
}

func WriteJsonPokemon(w http.ResponseWriter, pokemons []*data.Pokemon) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(pokemons)
}

func WriteJsonPokemonArray(w http.ResponseWriter, pokemons [][]interface{}) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(pokemons)
}
