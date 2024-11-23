package handler

import (
	"bayau/data"
	"bayau/util"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)

func PokemonHandler(w http.ResponseWriter, r *http.Request) {

	mapPokemon := make(map[int]bool)
	for len(mapPokemon) < 10 {
		randomNumber := rand.Intn(1024) + 1
		_, ok := mapPokemon[randomNumber]
		if !ok {
			mapPokemon[randomNumber] = true
		}
	}

	response := make([][]interface{}, 10)
	a := 0
	for i, _ := range mapPokemon {
		url := "https://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(i)
		resp, err := http.Get(url)
		if err != nil {
			util.WriteJsonError(w, err)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		pokemon := data.NewPokemon()
		err = json.Unmarshal(body, pokemon)
		if err != nil {
			util.WriteJsonError(w, err)
			return
		}
		response[a] = []interface{}{pokemon.Name, pokemon.Weight, pokemon.Sprite.ImageURL}
		a++
	}
	util.WriteJsonPokemonArray(w, response)
}
