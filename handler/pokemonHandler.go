package handler

import (
	"bayau/data"
	"bayau/db"
	"bayau/util"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

var NUMBER_POKEMONS = 1025
var POKEMONS_REQUESTED = 10

func PokemonMiddlewareHandler(db db.Persistance) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		mapRandomPokemon := util.RandomPokemons(POKEMONS_REQUESTED)
		chPokemons := make(chan *data.Pokemon)
		chPokemonsError := make(chan error)
		pokemons := make([]*data.Pokemon, 0, POKEMONS_REQUESTED)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
		defer cancel()
		start := time.Now()
		log.Println("Starting to fetch de pokemons...")

		for number := range mapRandomPokemon {
			go func() {
				url := "https://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(number)
				req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
				if err != nil {
					chPokemonsError <- err
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					chPokemonsError <- err
				}

				defer resp.Body.Close()
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					chPokemonsError <- err
				}
				pokemon := data.NewPokemon()
				err = json.Unmarshal(body, pokemon)
				if err != nil {
					chPokemonsError <- err
				}
				chPokemons <- pokemon
			}()
		}
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				util.WriteJsonError(w, ctx.Err(), http.StatusInternalServerError)
				return
			case err := <-chPokemonsError:
				util.WriteJsonError(w, err, http.StatusInternalServerError)
				return
			case pokemon := <-chPokemons:
				err := db.InsertPokemon(pokemon)
				if err != nil {
					util.WriteJsonError(w, err, http.StatusInternalServerError)
					return
				}
				pokemons = append(pokemons, pokemon)
			}
		}
		cancel() //this is secure
		log.Println("Pokemons fetched in", time.Since(start))
		util.WriteJsonPokemon(util.RESULT_OBJECT, w, pokemons)
	}
}
