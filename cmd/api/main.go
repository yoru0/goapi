package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/yoru0/goapi.git/internal/handlers"
	"github.com/yoru0/goapi.git/internal/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(middleware.Logger())
	r.Use(chiMiddleware.Recoverer)

	userHandler := handlers.NewUserHandler()

	r.Route("/api/v1", func(r chi.Router) {

		r.Route("/users", func(r chi.Router) {
			r.Post("/create", userHandler.UserCreate)
		})
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
