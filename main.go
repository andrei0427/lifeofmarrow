package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andrei0427/lifeofmarrow/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("error reading env file: %s", err))
	}

	initHttp()
}

func initHttp() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.Handle("/", handlers.HandleIndex())

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
