package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/oleiade/reflections"
	"io"
	"net/http"
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
		value, _ := reflections.GetField(resume, upcaseInitial(vars["1"]))
		log.Info(value)
		return value
	}

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

func AddNote(w http.ResponseWriter, r *http.Request) {

	//note := new(Notes)
	note := Note{"notes", "xyz", "xyz123"}
	//var storage LibKVBackend
	log.Info("Iniciate Storage backend")
	storage := NewLibKVBackend()

	log.Info(fmt.Sprintf("Add note [%s] and [%s]", note.Key, note.Note))
	storage.Put(note.Path, note.Key, []byte(note.Note))
	fmt.Fprint(w, "Welcome!\n")

}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	log.Info("Iniciate Storage backend")
	storage := NewLibKVBackend()
	note := Note{"notes", "", ""}

	log.Info(fmt.Sprintln("Get notes [%s]", note.Path))
	values := storage.GetAll(note.Path)
	json, err := json.Marshal(values)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(json))

}

func GetNote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	log.Info(vars["key"])
	log.Info("Iniciate Storage backend")
	storage := NewLibKVBackend()
	note := Note{"notes", "", ""}

	values := storage.Get(note.Path, vars["key"])
	json, err := json.Marshal(values)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(json))

}
