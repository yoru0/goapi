package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yoru0/goapi.git/internal/app/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/create", handlers.UserCreate)
			r.Post("/list", handlers.UserList)
			r.Post("/get", handlers.UserGet)
			r.Post("/update", handlers.UserUpdate)
		})
	})

	log.Println("Starting server on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
