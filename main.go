package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andrei0427/lifeofmarrow/internal"
	"github.com/andrei0427/lifeofmarrow/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("error reading env file: %s", err))
	}

	server := initHttp()
	log.Fatal(server.ListenAndServe())
}

func initHttp() *http.Server {
	mux := http.NewServeMux()
	internal.Cache = internal.NewCache()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/purge", handlers.HandlePurge)

	mux.HandleFunc("/about", handlers.HandleAbout)
	mux.HandleFunc("/recipes", handlers.HandleRecipes)

	mux.HandleFunc("/404", handlers.Handle404)
	mux.HandleFunc("/503", handlers.Handle503)

	mux.HandleFunc("/", handlers.HandleIndex)

	return &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
}
