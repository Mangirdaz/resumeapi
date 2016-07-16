package main

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"os"
)

func main() {

	log.Info("Starting JSON Resume backend")
	router := NewRouter()
	log.Fatal(http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), router))

}
