package main

import (
	"context"
	"encoding/csv"
	"flag"
	"interim_attestation/internal/city"
	"interim_attestation/internal/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gomodule/redigo/redis"
)

var (
	port     = flag.String("port", ":8080", "Port to listen on")
	addr     = flag.String("server", ":6379", "Redis server address")
	fileName = "cities.csv"
)

func init() {
	flag.Parse()
}

func main() {
	log.Println("Starting app")

	// Redis connection
	conn, err := redis.Dial("tcp", *addr)
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatalf("Failed to communicate to redis-server @ %v\n", err)
		}
	}()

	// Open file
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer CloseFile(file)

	// Read data from cvs-file to Redis
	reader := csv.NewReader(file)
	err = city.ReadInfo(conn, reader)
	if err != nil {
		log.Fatalln(err)
	}

	// The HTTP Server
	server := &http.Server{Addr: *port, Handler: service(conn)}

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

		// Write data from Redis back to cvs-file
		WriteChanges(conn, file)

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

func service(conn redis.Conn) http.Handler {
	log.Println("Starting server")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", handler.Get())

	r.Route("/city", func(r chi.Router) {
		r.Get("/{id}", handler.GetCity(conn))
	})

	r.Route("/region", func(r chi.Router) {
		r.Get("/{name}", handler.GetByRegion(conn))
	})

	r.Route("/district", func(r chi.Router) {
		r.Get("/{name}", handler.GetByDistrict(conn))
	})

	//the range parameter must contain 2 numbers separated by "-", for example "/100-3000"
	r.Route("/population", func(r chi.Router) {
		r.Get("/{range}", handler.GetByPopulation(conn))
	})

	//the range parameter must contain 2 numbers separated by "-", for example "/1480-1918"
	r.Route("/foundation", func(r chi.Router) {
		r.Get("/{range}", handler.GetByFoundation(conn))
	})

	r.Post("/create", handler.CreateCity(conn))
	r.Route("/update", func(r chi.Router) {
		r.Put("/{id}", handler.UpdatePopulation(conn))
	})

	r.Delete("/", handler.DeleteCity(conn))

	return r
}

func CloseFile(file *os.File) {
	log.Println("File closing")
	err := file.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func WriteChanges(conn redis.Conn, file *os.File) {
	log.Println("Writing changes to the file")
	err := city.WriteInfo(conn, file)
	if err != nil {
		log.Fatalln(err)
	}
}
