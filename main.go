package main

import (
	"api-todo/config"
	"api-todo/handlers"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/{id}", handlers.Get)
	r.Get("/", handlers.List)

	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)
}
