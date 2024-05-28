package main

import (
	"fmt"
	"go-short-url/config"
	"go-short-url/handlers"
	"log"
	"net/http"

	chi "github.com/go-chi/chi/v5"
)

func main() {

	config.ReadConfig()

	r := chi.NewRouter()

	// Routes
	r.Post("/short", handlers.ShortenURL)
	r.Get("/info", handlers.GetInfo)

	// Start the server
	address := fmt.Sprintf("%s:%d", config.Cfg.Service.Address, config.Cfg.Service.Port)
	log.Println("short url server started at", address)
	http.ListenAndServe(address, r)

}
