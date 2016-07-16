package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"
)

//Index index method for API
func Index(w http.ResponseWriter, r *http.Request, storage LibKVBackend) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "API")
}

//extend standart handler with our required storage backend details
type backendHandler func(w http.ResponseWriter, r *http.Request, storage LibKVBackend)

//retunr what mux expects
func mybackendHandler(handler backendHandler, storage LibKVBackend) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, storage)
	}
}

//Resume method to return JSON Resume
func Resume(w http.ResponseWriter, r *http.Request, storage LibKVBackend) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)

	var resume ResumeJson
	resume = readResume()
	var value interface{}
	log.Info("Var value ", vars)
	if len(vars) != 0 {
		log.Info("Get reseme part", vars)
		value = getResult(vars, resume)
	} else {
		log.Info("Full resume")
		value = resume
	}

	if err := json.NewEncoder(w).Encode(value); err != nil {
		panic(err)
	}
}

//AddNote add note to the key value storage
func AddNote(w http.ResponseWriter, r *http.Request, storage LibKVBackend) {

	//note := new(Notes)
	note := Note{"notes", "xyz", "xyz123"}
	//var storage LibKVBackend
	log.Info(fmt.Sprintf("Add note [%s] and [%s]", note.Key, note.Note))
	storage.Put(note.Path, note.Key, []byte(note.Note))
	fmt.Fprint(w, "Welcome!\n")
}

//GetNotes returns all notes
func GetNotes(w http.ResponseWriter, r *http.Request, storage LibKVBackend) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	note := Note{"notes", "", ""}

	log.Info(fmt.Sprintln("Get notes [%s]", note.Path))
	values, err := storage.GetAll(note.Path)
	if err != nil {
		fmt.Fprintln(w, "{'Error': 'Directory does not exist'}")
	} else {
		json, err := json.Marshal(values)
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(w, string(json))
	}

}

//GetNote get one note
func GetNote(w http.ResponseWriter, r *http.Request, storage LibKVBackend) {

	vars := mux.Vars(r)
	key := vars["key"]
	note := Note{"notes", "", ""}
	log.Info(key)
	if key != "" {
		log.Info("Key is not empty")
		values, err := storage.Get(note.Path, vars["key"])
		if err != nil {
			log.Error("Key not found")
			fmt.Fprintln(w, "{'Error': 'Key not found'}")
		} else {
			json, err := json.Marshal(values)
			if err != nil {
				log.Error("No key found")
				fmt.Fprintln(w, "{'Error': 'Key not found'}")
			} else {
				fmt.Fprintln(w, string(json))
			}
		}
	}
}
