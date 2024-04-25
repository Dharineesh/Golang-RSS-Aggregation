package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{}))

	v1Router := chi.NewRouter()

	v1Router.Get("/ready", handlerReadiness)

	v1Router.Get("/error", handlerReadinessErr)

	app := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	router.Mount("/v1", v1Router)

	log.Printf("Server is running and listening on the port: %v", portString)
	err := app.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
