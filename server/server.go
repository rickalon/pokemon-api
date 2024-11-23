package server

import (
	"bayau/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run(addr string) {
	router := mux.NewRouter()

	router.HandleFunc("/pokemon", handler.PokemonHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(addr, router))
}
