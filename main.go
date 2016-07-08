package main

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
)

func main() {

	router := NewRouter()

	go func() {
		log.Fatal(http.ListenAndServe("localhost:8081", router))
	}()
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
