package util

import "math/rand"

func RandomPokemons(max int) map[int]int {
	counter := 0
	mapPokemon := make(map[int]int)
	for len(mapPokemon) < max {
		randomNumber := rand.Intn(max) + 1
		_, ok := mapPokemon[randomNumber]
		if !ok {
			mapPokemon[randomNumber] = counter
			counter++
		}
	}
	return mapPokemon
}
