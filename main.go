package main

import (
	"fmt"
	"go-short-url/configs"
	"go-short-url/handlers"
	"go-short-url/services"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	chi "github.com/go-chi/chi/v5"
)

func main() {

	configs.ReadConfig()

	router := chi.NewRouter()
	// Middleware setup
	router.Use(middleware.Logger)

	// Routes
	urlSvc := services.NewShortenURLServicer()
	router.Mount("/v1/urlshorter", handlers.NewShortenURLHandler(urlSvc))

	// Start the server
	address := fmt.Sprintf("%s:%d", configs.Cfg.Service.Address, configs.Cfg.Service.Port)
	log.Println("short url server started at", address)
	http.ListenAndServe(address, router)

}
