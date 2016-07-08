package main

import (
	//"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	//"strconv"
	"bytes"
	"encoding/json"
	"github.com/oleiade/reflections"
	"reflect"
	"strings"
	"unicode"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)

	log.Info(reflect.TypeOf(vars))

	//TODO : local caching as it will do uneeded calls to static content
	var resume ResumeJson
	resume = readResume()
	value := getResult(vars, resume)

	if err := json.NewEncoder(w).Encode(value); err != nil {
		panic(err)
	}
}

func getResult(vars map[string]string, resume ResumeJson) interface{} {
	var value interface{}
	if len(vars) == 1 {
		value, err := reflections.GetField(resume, upcaseInitial(vars["1"]))
		if err != nil {
			panic(err)
		}
		return value
	}

	//value, err := reflections.GetField(resume, upcaseInitial(vars["1"]))

	return value
}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

func upcaseInitial(str string) string {
	for i, v := range str {
		var end string
		end = str[i+1:]
		return string(unicode.ToUpper(v)) + strings.ToLower(end)
	}
	return ""
}
