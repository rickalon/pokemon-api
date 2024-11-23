package server

import (
	"bayau/handler"
	"bayau/settings"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server interface {
	Run(cfg *settings.Config)
}

type DefaultServer struct{}

func Run(srv server, cfg *settings.Config) {
	srv.Run(cfg)
}

func NewDefaultServer() *DefaultServer {
	return &DefaultServer{}
}

func (serv *DefaultServer) Run(cfg *settings.Config) {
	log.Println("Configuring the server.", cfg)
	router := mux.NewRouter()
	log.Println("Added handler pokemons")
	router.HandleFunc("/pokemons", handler.PokemonHandler).Methods("GET")
	log.Println("Server running")
	log.Fatal(http.ListenAndServe(cfg.Addr, router))
}
