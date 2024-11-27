package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joaxaca/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	fmt.Println(`
	________  ___       ________  ________   ___  __            ________  ________  ___          ________  ________  ___       
	|\   __  \|\  \     |\   __  \|\   ___  \|\  \|\  \         |\   __  \|\   __  \|\  \        |\   ____\|\   __  \|\  \      
	\ \  \|\  \ \  \    \ \  \|\  \ \  \\ \  \ \  \/  /|_       \ \  \|\  \ \  \|\  \ \  \       \ \  \___|\ \  \|\  \ \  \     
	 \ \   ____\ \  \    \ \   __  \ \  \\ \  \ \   ___  \       \ \   __  \ \   ____\ \  \       \ \  \  __\ \  \\\  \ \  \    
	  \ \  \___|\ \  \____\ \  \ \  \ \  \\ \  \ \  \\ \  \       \ \  \ \  \ \  \___|\ \  \       \ \  \|\  \ \  \\\  \ \__\   
	   \ \__\    \ \_______\ \__\ \__\ \__\\ \__\ \__\\ \__\       \ \__\ \__\ \__\    \ \__\       \ \_______\ \_______\|__|   
	    \|__|     \|_______|\|__|\|__|\|__| \|__|\|__| \|__|        \|__|\|__|\|__|     \|__|        \|_______|\|_______|   ___ 
	                                                                                                                       |\__\
	                                                                                                                       \|__|
	`)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}