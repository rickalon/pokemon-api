package util

import "math/rand"

func RandomPokemons(max int) map[int]interface{} {
	mapPokemon := make(map[int]interface{})
	for len(mapPokemon) < max {
		randomNumber := rand.Intn(max) + 1
		_, ok := mapPokemon[randomNumber]
		if !ok {
			mapPokemon[randomNumber] = 1
		}
	}
	return mapPokemon
}
