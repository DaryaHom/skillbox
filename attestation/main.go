package main

import (
	"attestation/internal/str"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jszwec/csvutil"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type alphaCode struct {
	Name string `csv:"Name"`
	Code string `csv:"Code"`
}

var (
	alphaCodes    map[string]string
	host          = "http://localhost"
	simulatorAddr = ":8383"
	addr          = ":8585"
)

func init() {

	// Get alpha-2 country codes data from cvs file
	alphaCodes = make(map[string]string)

	file, err := os.OpenFile("./assets/ISO3166-1alpha-2.csv", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	dec, err := csvutil.NewDecoder(reader)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		a := alphaCode{}

		if err := dec.Decode(&a); err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}
		alphaCodes[a.Code] = a.Name
	}

	// Run simulator
	err = runSimulator()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// A new router
	r := mux.NewRouter()

	// Testing connection handler
	r.HandleFunc("/", handleConnection)

	// A handler to get all service data
	r.HandleFunc("/struct", handleData)

	// Server
	log.Println("Listening on", addr)
	log.Fatalln(http.ListenAndServe(addr, r))
}

// handleConnection - checks connection
func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", "ok")
}

// handleData - returns structured data from all services
func handleData(w http.ResponseWriter, r *http.Request) {
	resultT := str.GetResultData(alphaCodes, host, simulatorAddr)

	res, err := json.Marshal(*resultT)
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := w.Write(res); err != nil {
		log.Fatalln(err)
	}
}

func runSimulator() error {
	cmd := exec.Command("./build/simulator.exe")
	var buf bytes.Buffer
	cmd.Stdout = &buf
	err := cmd.Start()
	if err != nil {
		return err
	}
	return nil
}
