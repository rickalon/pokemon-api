package db

import "bayau/data"

type Persistance interface {
	InsertPokemon(pokemon *data.Pokemon) error
	InsertPokemons(pokemon ...*data.Pokemon) error
}
