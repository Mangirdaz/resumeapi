package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func readResume() ResumeJson {

	response, err := http.Get("https://raw.githubusercontent.com/Mangirdaz/resume/master/Mangirdas.Judeikis.json")
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
	}

	var resume ResumeJson
	hah, err := ioutil.ReadAll(response.Body)
	error := json.Unmarshal(hah, &resume)
	if err != nil {
		log.Fatal(error)
	}
	return resume

}
