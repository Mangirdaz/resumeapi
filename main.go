package main

import (
	log "github.com/Sirupsen/logrus"
)

func main() {

	log.Info("Starting JSON Resume backend")
	NewRouter()
	//log.Fatal(http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), router))

}
