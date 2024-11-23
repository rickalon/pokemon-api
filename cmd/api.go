package main

import (
	"bayau/db"
	"bayau/server"
	"bayau/settings"
)

func main() {
	cfg := settings.NewConfig(settings.ConfigAddr(":8080"))
	srv := server.NewDefaultServer()
	db := db.NewMongoDB()
	server.Run(srv, cfg, db)
}
