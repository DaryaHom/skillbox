package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"proxy/internal/req"

	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

const proxyAddr string = ":9000"

var (
	counter            = 0
	firstInstanceHost  = "http://localhost:8080"
	secondInstanceHost = "http://localhost:8081"
)

func main() {
	log.Printf("Listening on %s, forwarding to...\n", proxyAddr)
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/", handleGet)
		r.Put("/{user_id}", handlePut)
	})
	r.Route("/friends", func(r chi.Router) {
		r.Get("/{user_id}", handleGet)
	})
	r.Post("/create", handlePost)
	r.Post("/make_friends", handlePost)
	r.Delete("/user", handleDel)

	log.Fatalln(http.ListenAndServe(proxyAddr, r))
}

//handleGet - sends get-request to the instance
func handleGet(w http.ResponseWriter, r *http.Request) {
	host := fmt.Sprintf("%s%v", chooseInstanceHost(), r.URL)
	log.Println(host)
	resp, err := http.Get(host)
	if err != nil {
		log.Fatalln(err)
	}

	textBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.CloseRequestBody(resp.Body)

	if _, err := w.Write(textBytes); err != nil {
		log.Fatalln(err)
	}
}

//handlePost - sends post-request to the instance
func handlePost(w http.ResponseWriter, r *http.Request) {
	host := fmt.Sprintf("%s%v", chooseInstanceHost(), r.URL)
	log.Println(host)

	resp, err := http.Post(host, "application/json", r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.CloseRequestBody(r.Body)

	textBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.CloseRequestBody(resp.Body)

	if _, err := w.Write(textBytes); err != nil {
		log.Fatalln(err)
	}
}

//handleDel - sends delete-request to the instance
func handleDel(w http.ResponseWriter, r *http.Request) {
	host := fmt.Sprintf("%s%v", chooseInstanceHost(), r.URL)
	log.Println(host)

	request, err := http.NewRequest("DELETE", host, r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.CloseRequestBody(r.Body)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	textBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.CloseRequestBody(resp.Body)

	if _, err := w.Write(textBytes); err != nil {
		log.Fatalln(err)
	}
}

//handlePut - sends put-request to the instance
func handlePut(w http.ResponseWriter, r *http.Request) {
	host := fmt.Sprintf("%s%v", chooseInstanceHost(), r.URL)
	log.Println(host)

	request, err := http.NewRequest("PUT", host, r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.CloseRequestBody(r.Body)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	textBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.CloseRequestBody(resp.Body)

	if _, err := w.Write(textBytes); err != nil {
		log.Fatalln(err)
	}
}

func chooseInstanceHost() (instance string) {
	if counter == 0 {
		instance = firstInstanceHost
		counter++
		return
	}

	instance = secondInstanceHost
	counter--
	return
}
