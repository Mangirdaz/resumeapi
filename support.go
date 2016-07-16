package main

import (
	"bytes"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/oleiade/reflections"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"unicode"
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

//get results from json resume
func getResult(vars map[string]string, resume ResumeJson) interface{} {
	var value interface{}
	if len(vars) == 1 {
		value, _ := reflections.GetField(resume, upcaseInitial(vars["1"]))
		return value
	}
	return value
}

//StreamToByte stream to bytes conversion
func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

//Becase we call data from struct we are key sensitive. So we make this hack. TODO: migrate to annotations
func upcaseInitial(str string) string {
	for i, v := range str {
		var end string
		end = str[i+1:]
		return string(unicode.ToUpper(v)) + strings.ToLower(end)
	}
	return ""
}
