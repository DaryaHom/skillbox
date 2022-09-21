package main

import (
	"context"
	"encoding/csv"
	"flag"
	"interim_attestation/internal/handler"
	"interim_attestation/internal/storage"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	port     = flag.String("port", ":8080", "port to listen on")
	fileName = "cities.csv"
	s        = storage.NewStorage()
)

func init() {
	flag.Parse()
}

func main() {
	// Opening file
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer CloseFile(file)

	reader := csv.NewReader(file)
	err = s.ReadCitiesInfo(reader)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Starting app")

	// The HTTP Server
	server := &http.Server{Addr: *port, Handler: service()}

	// Server run context
	ctx, cancelCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(ctx, 30*time.Second)
		log.Println("Interrupt signal received")
		WriteChanges(s, file)

		go func() {
			<-shutdownCtx.Done()
			log.Println("Completion of work")

			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("Graceful shutdown timed out.. forcing exit")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		log.Println("Graceful shutdown")
		if err != nil {
			log.Fatal(err)
		}
		cancelCtx()
	}()

	// Run the server
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-ctx.Done()
}

func service() http.Handler {
	log.Println("Starting server")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", handler.Get())
	r.Route("/city", func(r chi.Router) {
		r.Get("/{id}", handler.GetCity(s))
	})

	r.Route("/region", func(r chi.Router) {
		r.Get("/{name}", handler.GetByRegion(s))
	})

	r.Route("/district", func(r chi.Router) {
		r.Get("/{name}", handler.GetByDistrict(s))
	})

	//the range parameter must contain 2 numbers separated by "-", for example "/100-3000"
	r.Route("/population", func(r chi.Router) {
		r.Get("/{range}", handler.GetByPopulation(s))
	})

	//the range parameter must contain 2 numbers separated by "-", for example "/1480-1918"
	r.Route("/foundation", func(r chi.Router) {
		r.Get("/{range}", handler.GetByFoundation(s))
	})

	r.Post("/create", handler.CreateCity(s))
	r.Route("/update", func(r chi.Router) {
		r.Put("/{id}", handler.UpdatePopulation(s))
	})

	r.Delete("/", handler.DeleteCity(s))

	return r
}

func CloseFile(file *os.File) {
	log.Println("File closing")
	err := file.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func WriteChanges(s *storage.Storage, file *os.File) {
	log.Println("Writing changes to the file")
	err := s.WriteCitiesInfo(file)
	if err != nil {
		log.Fatalln(err)
	}
}
