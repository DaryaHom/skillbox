package main

import (
	"log"
	"net/http"
	"network_communication/internal/handler"
	"network_communication/pkg/storage"

	"github.com/go-chi/chi"
)

var (
	port = ":8080"
)

func main() {
	r := NewHTTPRouter()

	log.Println("Serving on" + port)
	log.Fatal(http.ListenAndServe(port, r))
}

// NewHTTPRouter - returns new chi-router
func NewHTTPRouter() *chi.Mux {
	r := chi.NewRouter()
	s := storage.NewStorage()

	r.Get("/getAll", handler.GetAll(s))
	r.Post("/create", handler.CreateUser(s))
	r.Post("/make_friends", handler.MakeFriends(s))
	r.Delete("/user", handler.DeleteUser(s))

	r.Route("/friends", func(r chi.Router) {
		r.Get("/{user_id}", handler.GetAllFriends(s))
	})

	r.Route("/", func(r chi.Router) {
		r.Get("/", handler.Get())
		r.Put("/{user_id}", handler.UpdateAge(s))
	})

	return r
}
