package req

import (
	"io"
	"log"
	"net/http"
)

func CloseRequestBody(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

//WriteOkResponse - returns 200-status
func WriteOkResponse(w http.ResponseWriter, response []byte) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		log.Fatalln(err)
	}
}

//WriteCreateResponse - returns 201-status
func WriteCreateResponse(w http.ResponseWriter, response []byte) {
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(response); err != nil {
		log.Fatalln(err)
	}
}

//WriteBadRequest - returns 400-status
func WriteBadRequest(w http.ResponseWriter, response []byte) {
	w.WriteHeader(http.StatusBadRequest)
	if _, err := w.Write(response); err != nil {
		log.Fatalln(err)
	}
}
