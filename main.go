package main

import (
	loggo "github.com/juju/loggo"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error while opening file %v\n", err)
	}

	log.SetOutput(file)

	for i := 0; i < 10; i++ {
		log.Println("Rrunning...")
		loggo.LoggerInfo()
	}
}
