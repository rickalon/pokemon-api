package db

import (
	"bayau/data"
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

type MongoDB struct {
	session *mgo.Session
}

func defaultMongo() *MongoDB {
	return &MongoDB{}
}

func NewMongoDB() *MongoDB {
	dflt := defaultMongo()
	session, err := mgo.Dial(os.Getenv("ME_CONFIG_MONGODB_URL"))
	if err != nil {
		log.Fatalf("Failing to start a mongoDB session, %s", err)
	}
	dflt.session = session
	return dflt
}

func (m *MongoDB) InsertPokemon(pokemon *data.Pokemon) error {
	session := m.session.Copy()
	defer session.Close()
	collection := session.DB("Items").C("Pokemons")
	return collection.Insert(pokemon)
}

func (m *MongoDB) InsertPokemons(pokemons ...*data.Pokemon) error {
	session := m.session.Copy()
	defer session.Close()
	collection := m.session.DB("Items").C("Pokemons")
	return collection.Insert(pokemons)
}
