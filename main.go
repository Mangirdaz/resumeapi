package main

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"os"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("PORT"), router))

}
