package service

import (
	"interim_attestation/internal/db"
	"interim_attestation/internal/handler"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Service(redisDB *db.DB) http.Handler {
	log.Println("Starting server")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", handler.Get())

	r.Route("/city", func(r chi.Router) {
		r.Get("/{id}", handler.GetCity(redisDB))
	})

	r.Route("/region", func(r chi.Router) {
		r.Get("/{name}", handler.GetByRegion(redisDB))
	})

	r.Route("/district", func(r chi.Router) {
		r.Get("/{name}", handler.GetByDistrict(redisDB))
	})

	//the range parameter must contain 2 numbers separated by "-", for example "/100-3000"
	r.Route("/population", func(r chi.Router) {
		r.Get("/{range}", handler.GetByPopulation(redisDB))
	})

	//the range parameter must contain 2 numbers separated by "-", for example "/1480-1918"
	r.Route("/foundation", func(r chi.Router) {
		r.Get("/{range}", handler.GetByFoundation(redisDB))
	})

	r.Post("/create", handler.CreateCity(redisDB))
	r.Route("/update", func(r chi.Router) {
		r.Put("/{id}", handler.UpdatePopulation(redisDB))
	})

	r.Delete("/", handler.DeleteCity(redisDB))

	return r
}
