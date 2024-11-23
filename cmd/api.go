package main

import (
	"bayau/server"
	"bayau/settings"
)

func main() {
	cfg := settings.NewConfig(settings.ConfigAddr(":8080"))
	srv := server.NewDefaultServer()
	server.Run(srv, cfg)
}
