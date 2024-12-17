package main

import (
	"net/http"
	"github.com/joaxaca/goapi/internal/handlers"
	"github.com/joaxaca/goapi/internal/middleware"
	"log"
	"os"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/account/coins", handlers.GetCoinBalance)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: middleware.StripSlashes(middleware.Authorization(router)),
	}

	log.Println("Starting GO API service on port :" + port + "...")
	server.ListenAndServe()
}