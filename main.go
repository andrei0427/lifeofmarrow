package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andrei0427/lifeofmarrow/internal"
	"github.com/andrei0427/lifeofmarrow/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("error reading env file: %s", err))
	}

	log.Fatal(initHttp())
}

func initHttp() error {
	mux := http.NewServeMux()
	internal.Cache = internal.NewCache()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/purge", handlers.HandlePurge)
	mux.HandleFunc("/about", handlers.HandleAbout)
	mux.HandleFunc("/recipes/{p...}", handlers.HandleRecipesPage)
	mux.HandleFunc("/recipe/{slug}", handlers.HandleRecipePage)
	mux.HandleFunc("/store/books", handlers.HandleBooks)

	mux.HandleFunc("/404", handlers.Handle404)
	mux.HandleFunc("/503", handlers.Handle503)

	mux.HandleFunc("/", handlers.HandleIndex)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	return http.ListenAndServe(fmt.Sprintf("localhost:%s", port), mux)
}
