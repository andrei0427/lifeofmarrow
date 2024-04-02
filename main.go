package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andrei0427/lifeofmarrow/internal"
	"github.com/andrei0427/lifeofmarrow/internal/handlers"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/stripe/stripe-go/v76"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("error reading env file: %s", err))
	}

	key, ok := os.LookupEnv("STRIPE_KEY")
	if !ok {
		log.Fatalf("Stripe API key not found in env")
	}

	stripe.Key = key

	log.Fatal(initHttp())
}

func initHttp() error {
	mux := http.NewServeMux()
	internal.Cache = internal.NewCache()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/purge", handlers.HandlePurge)
	mux.HandleFunc("/about", handlers.HandleAbout)
	mux.HandleFunc("/contact", handlers.HandleContact)

	mux.HandleFunc("/recipes/{p...}", handlers.HandleRecipesPage)
	mux.HandleFunc("/recipe/{slug}", handlers.HandleRecipePage)

	mux.HandleFunc("GET /store/books", handlers.HandleBooks)
	mux.HandleFunc("GET /store/food", handlers.HandleFood)
	mux.HandleFunc("POST /store/checkout/{id}", handlers.HandleCheckout)
	mux.HandleFunc("GET /store/checkout/cancel", handlers.HandleCheckoutCancel)
	mux.HandleFunc("GET /store/checkout/success", handlers.HandleCheckoutSuccess)
	mux.HandleFunc("POST /store/webhook", handlers.HandleStripeWebhook)

	mux.HandleFunc("/404", handlers.Handle404)
	mux.HandleFunc("/503", handlers.Handle503)

	mux.HandleFunc("/", handlers.HandleIndex)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	corsMux := cors.Default().Handler(mux)
	return http.ListenAndServe(fmt.Sprintf("localhost:%s", port), corsMux)
}
