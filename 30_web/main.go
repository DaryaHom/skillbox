package main

import (
	"log"
	"net/http"
	"network_communication/internal/handler"
	"network_communication/pkg/storage"

	"github.com/go-chi/chi"
)

func main() {
	port := ":8080"
	r := chi.NewRouter()
	s := storage.NewStorage()

	log.Println("Starting server")

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

	log.Println("Serving on" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
