package main

import (
	"net/http"
	"github.com/joaxaca/goapi/internal/handlers"
	"github.com/joaxaca/goapi/internal/middleware"
	"log"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/account/coins", handlers.GetCoinBalance)

	server := &http.Server{
		Addr:    ":8000",
		Handler: middleware.StripSlashes(middleware.Authorization(router)),
	}

	log.Println("Starting GO API service on port :8000...")
	server.ListenAndServe()
}