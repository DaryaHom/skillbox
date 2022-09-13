package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"proxy/internal/req"

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
	http.HandleFunc("/", handle)
	log.Fatalln(http.ListenAndServe(proxyAddr, nil))
}

//handlePut - sends put-request to the instance
func handle(w http.ResponseWriter, r *http.Request) {
	host := fmt.Sprintf("%s%v", chooseInstanceHost(), r.URL)
	log.Println(host)

	request, err := http.NewRequest(r.Method, host, r.Body)
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
