package server

import (
	"bayau/db"
	"bayau/handler"
	"bayau/settings"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server interface {
	Run(cfg *settings.Config, db db.Persistance)
}

type DefaultServer struct{}

func Run(srv server, cfg *settings.Config, db db.Persistance) {
	srv.Run(cfg, db)
}

func NewDefaultServer() *DefaultServer {
	return &DefaultServer{}
}

func (serv *DefaultServer) Run(cfg *settings.Config, db db.Persistance) {
	log.Println("Configuring the server.", cfg)
	router := mux.NewRouter()
	log.Println("Added handler pokemons")
	router.HandleFunc("/pokemons", handler.PokemonMiddlewareHandler(db)).Methods("GET")
	log.Println("Server running")
	log.Fatal(http.ListenAndServe(cfg.Addr, router))
}
