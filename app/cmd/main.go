package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"github.com/kingsbloc/scissor/internal/config"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	port := os.Getenv("PORT")

	// Init the main Router
	r := chi.NewRouter()

	// Add Middlewares
	r.Use(middleware.CleanPath)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Connect Redis
	config.ConnectRedis()

	// Serve and listen
	log.Fatal(http.ListenAndServe(":"+port, r))
}
