package main

import (
	"fmt"
	"go-short-url/configs"
	"go-short-url/handlers"
	"go-short-url/services"
	"log"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// load configs
	configs.ReadConfig()

	// initializing services
	urlSvc := services.NewShortenURLServicer()

	// creating router
	router := createRouter(urlSvc)

	// Start the server
	address := fmt.Sprintf("%s:%d", configs.Cfg.Service.Address, configs.Cfg.Service.Port)

	log.Println("short url server started at", address)
	http.ListenAndServe(address, router)
}

func createRouter(urlSvc services.ShortenURLServicer) chi.Router {
	router := chi.NewRouter()
	// Middleware setup
	router.Use(middleware.Logger)

	// Routes
	router.Mount("/v1/urlshorter", handlers.NewShortenURLHandler(urlSvc))

	return router
}
