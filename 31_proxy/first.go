package main

import (
	"database/sql"
	"log"
	"net/http"
	"proxy/internal/handler"

	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	port := ":8080"
	r := chi.NewRouter()

	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		log.Fatalln(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln()
		}
	}(db)

	log.Println("Starting server 1")

	r.Route("/", func(r chi.Router) {
		r.Get("/", handler.Get())
		r.Put("/{user_id}", handler.UpdateAge(db))
	})
	r.Post("/create", handler.CreateUser(db))
	r.Post("/make_friends", handler.MakeFriends(db))
	r.Delete("/user", handler.DeleteUser(db))

	r.Route("/friends", func(r chi.Router) {
		r.Get("/{user_id}", handler.GetAllFriends(db))
	})

	log.Println("Serving on" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
