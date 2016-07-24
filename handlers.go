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

type Handler interface {
	ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

//retunr what mux expects
func mybackendHandler(handler backendHandler, storage LibKVBackend) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, storage)
	}
}
func CheckAuth(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	user, pass, _ := r.BasicAuth()
	if !checkPass(user, pass) {
		printStatus(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	next(w, r)
}

func checkPass(user, pass string) bool {
	log.Info(fmt.Sprintf("User [%s] and Pass [%s]", user, pass))
	if user == "mj" && pass == "test" {
		log.Info("Pass OK")
		return true
	} else {
		log.Info("Pass Error")
		return false
	}
	return false
}

//Resume method to return JSON Resume
func Resume(w http.ResponseWriter, r *http.Request, storage LibKVBackend) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)

	resume := readResume()
	var value interface{}
	log.Info(fmt.Sprintf("Vars in payload [%s]", vars))
	if len(vars) != 0 {
		log.Info(fmt.Sprintf("Reseme part [%s]", vars))
		value = getResult(vars, resume)
	} else {
		log.Info("Full resume request")
		value = resume
	}

	if err := json.NewEncoder(w).Encode(value); err != nil {
		panic(err)
	}
}

//AddNote add note to the key value storage
func AddNotes(w http.ResponseWriter, r *http.Request, storage LibKVBackend) {
	log.Info("Add Notes request")
	//TODO: Add array post option
	decoder := json.NewDecoder(r.Body)
	var note Note
	err := decoder.Decode(&note)
	if err != nil {
		log.Error(fmt.Sprintf("Error while parsing payload [%s]", err))
		printStatus(w, http.StatusExpectationFailed, "Error while parsing payload ")
	} else if note.Key == "" || note.Note == "" {
		log.Error(fmt.Sprintf("Error while parsing payload. Value missing "))
		printStatus(w, http.StatusExpectationFailed, "Error while parsing payload ")
	} else {
		//var storage LibKVBackend
		log.Info(fmt.Sprintf("Add note [%s] [%s]", note.Note, note.Key))
		note.Path = "notes"
		err = storage.Put(note.Path, note.Key, []byte(note.Note))
		if err == nil {
			printStatus(w, http.StatusOK, "Note Stored")
		} else {
			log.Error(fmt.Sprintf("Note not stored with error [%s]", err))
			printStatus(w, http.StatusNotImplemented, "Note Not Stored")
		}
	}

}

//GetNotes returns all notes
func GetNotes(w http.ResponseWriter, r *http.Request, storage LibKVBackend) {
	var note Note
	note.Path = "notes"

	log.Info(fmt.Sprintf("Get notes [%s]", note.Path))
	values, err := storage.GetAll(note.Path)
	if err != nil {
		log.Error(fmt.Sprintf("Directory [%s] not found", note.Path))
		printStatus(w, http.StatusNoContent, "Directory not found")
	} else {
		json, err := json.Marshal(values)
		if err != nil {
			log.Error(fmt.Sprintf("Note not recieved with error [%s]", err))
			printStatus(w, http.StatusNoContent, "Error while getting values")
		}
		fmt.Fprintln(w, string(json))
	}

}

//GetNote get one note
func GetNote(w http.ResponseWriter, r *http.Request, storage LibKVBackend) {

	vars := mux.Vars(r)
	key := vars["key"]
	var note Note
	note.Path = "notes"
	log.Info(fmt.Sprintf("Looking for key [%s]", key))
	if key != "" {
		log.Info("Key is not empty")
		values, err := storage.Get(note.Path, vars["key"])
		if err != nil {
			log.Error(fmt.Sprintf("Key [%s] not found", key))
			printStatus(w, http.StatusNoContent, "Key not found")
		} else {
			json, err := json.Marshal(values)
			if err != nil {
				log.Error("No key found")
				printStatus(w, http.StatusNoContent, "Key not found")
			} else {
				fmt.Fprintln(w, string(json))

			}
		}
	}
}

func printStatus(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	response := ErrorMessage{status, message}
	json, err := json.Marshal(response)
	log.Info("Message to return: " + string(json))
	log.Info(err)
	if err == nil {
		fmt.Fprintln(w, string(json))
	}

}
